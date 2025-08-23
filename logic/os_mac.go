//go:build darwin

package logic

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const osaTimeout = 5 * time.Second

// for env vars like $HOME
func expandToAbsPath(p string) (string, error) {
	if p == "" {
		return "", errors.New("empty path")
	}
	p = filepath.Clean(p)

	if strings.Contains(p, "$") {
		p = os.ExpandEnv(p)
	}

	//change ~ to home dir
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

// timeout for the osascript helper func
func runOSA(ctx context.Context, script string) (string, error) {
	if _, err := exec.LookPath("osascript"); err != nil {
		return "", fmt.Errorf("osascript not found: %w", err)
	}
	cmd := exec.CommandContext(ctx, "osascript", "-e", script)
	out, err := cmd.Output()
	if ctx.Err() != nil {
		return "", fmt.Errorf("osascript timeout: %w", ctx.Err())
	}
	if err != nil {
		return "", fmt.Errorf("osascript failed: %w", err)
	}
	return string(out), nil
}

// NOTE: this only get the main desktop wallpaper
func Get() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//using System Events
	if out, err := runOSA(ctx, `tell application "System Events" to get POSIX path of (picture of desktop 1)`); err == nil {
		if s := strings.TrimSpace(out); s != "" {
			return s, nil
		}
	}

	// using Finder
	out, err := runOSA(ctx, `tell application "Finder" to get POSIX path of (get desktop picture as alias)`)
	if err != nil {
		return "", fmt.Errorf("get wallpaper via System Events/Finder failed: %w", err)
	}
	return strings.TrimSpace(out), nil
}

func SetFromFile(path string) error {
	abs, err := expandToAbsPath(path)
	if err != nil {
		return err
	}
	if _, err := os.Stat(abs); err != nil {
		return fmt.Errorf("stat: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), osaTimeout)
	defer cancel()

	script := `tell application "System Events" to tell every desktop to set picture to POSIX file ` + strconv.Quote(abs)
	if _, err := runOSA(ctx, script); err != nil {
		return err
	}
	return nil
}

// INFO: not working on macOS
type Mode int

func Set(path string, _ Mode) error {
	return SetFromFile(path)
}

func getCacheDir() (string, error) {
	dir, err := os.UserCacheDir() // ~/Library/Caches
	if err != nil {
		return "", fmt.Errorf("cache dir: %w", err)
	}
	// make sure the dir exists
	if _, statErr := os.Stat(dir); os.IsNotExist(statErr) {
		if mkErr := os.MkdirAll(dir, 0o755); mkErr != nil {
			return "", fmt.Errorf("mkdir cache dir: %w", mkErr)
		}
	}
	return dir, nil
}
