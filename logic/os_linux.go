//go:build linux

package logic

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const waypaperTimeout = 5 * time.Second

type Mode int

// expandToAbsPath expands user-friendly paths to a clean absolute path.
// It supports:
//   - Environment variables like $HOME
//   - Tilde expansion: ~ and ~/something
//   - Normalization (removing ./, ../ where possible)
func expandToAbsPath(p string) (string, error) {
	if p == "" {
		return "", errors.New("empty path")
	}
	p = filepath.Clean(p)

	if strings.Contains(p, "$") {
		p = os.ExpandEnv(p)
	}

	if strings.HasPrefix(p, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("home: %w", err)
		}
		if p == "~" {
			p = home
		} else if strings.HasPrefix(p, "~/") {
			p = filepath.Join(home, p[2:])
		}
	}

	abs, err := filepath.Abs(p)
	if err != nil {
		return "", fmt.Errorf("abs: %w", err)
	}
	return abs, nil
}

func runCmdWithTimeout(timeout time.Duration, name string, args ...string) (string, error) {
	if _, err := exec.LookPath(name); err != nil {
		return "", fmt.Errorf("%s not found in PATH: %w", name, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	out, err := cmd.CombinedOutput()

	if ctx.Err() != nil {
		return "", fmt.Errorf("%s timeout: %w", name, ctx.Err())
	}
	if err != nil {
		return "", fmt.Errorf("%s failed: %w (output: %s)", name, err, strings.TrimSpace(string(out)))
	}
	return string(out), nil
}

func mustReadableFile(p string) error {
	info, err := os.Stat(p)
	if err != nil {
		return fmt.Errorf("stat: %w", err)
	}
	if !info.Mode().IsRegular() {
		return fmt.Errorf("not a regular file: %s", p)
	}
	f, err := os.Open(p)
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	_ = f.Close()
	return nil
}

func userCacheFile() (string, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", fmt.Errorf("cache dir: %w", err)
	}
	appDir := filepath.Join(dir, "waraqah")
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		if mkErr := os.MkdirAll(appDir, 0o755); mkErr != nil {
			return "", fmt.Errorf("mkdir cache dir: %w", mkErr)
		}
	}
	return filepath.Join(appDir, "current_wallpaper"), nil
}

func writeCurrentToCache(abs string) {
	if fpath, err := userCacheFile(); err == nil {
		_ = os.WriteFile(fpath, []byte(abs+"\n"), 0o644)
	}
}

func readCurrentFromCache() (string, error) {
	fpath, err := userCacheFile()
	if err != nil {
		return "", err
	}
	b, err := os.ReadFile(fpath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}

// readFromWaypaperConfig tries to read the current wallpaper from Waypaper config:
// NOTE: is this the path ~/.config/waypaper/config.ini
// It looks for a line like: wallpaper = /path/to/image.jpg
func readFromWaypaperConfig() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	cfg := filepath.Join(home, ".config", "waypaper", "config.ini")
	f, err := os.Open(cfg)
	if err != nil {
		return "", err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}
		if strings.HasPrefix(strings.ToLower(line), "wallpaper") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				val := strings.TrimSpace(parts[1])
				if val != "" {
					return val, nil
				}
			}
		}
	}
	if err := sc.Err(); err != nil {
		return "", err
	}
	return "", errors.New("wallpaper not found in waypaper config")
}

// Get returns the current wallpaper path.
// It tries two strategies:
//  1. Read from Waypaper config (if present)
//  2. Fallback to our local cache if we were the last to set it
func Get() (string, error) {
	if s, err := readFromWaypaperConfig(); err == nil && s != "" {
		return s, nil
	}
	if s, err := readCurrentFromCache(); err == nil && s != "" {
		return s, nil
	}
	return "", errors.New("unable to determine current wallpaper (waypaper config/cache not found)")
}

func SetFromFile(path string) error {
	abs, err := expandToAbsPath(path)
	if err != nil {
		return err
	}
	if err := mustReadableFile(abs); err != nil {
		return err
	}

	_, err = runCmdWithTimeout(waypaperTimeout, "waypaper", "--wallpaper", abs)
	if err != nil {
		return err
	}

	writeCurrentToCache(abs)
	return nil
}

// On Linux, lik mac 'mode' is ignored.
func Set(path string, _ Mode) error {
	return SetFromFile(path)
}
