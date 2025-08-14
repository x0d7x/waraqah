//go:build windows

package logic

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf16"
	"unsafe"

	// "syscall" changed to "golang.org/x/sys/windows"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

const (
	spiGetDeskWallpaper = 0x0073
	spiSetDeskWallpaper = 0x0014
	spifUpdateINIFile   = 0x01
	spifSendChange      = 0x02
)

var (
	user32  = windows.NewLazySystemDLL("user32.dll")
	spiProc = user32.NewProc("SystemParametersInfoW")
)

// NOTE: https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-systemparametersinfow

const maxPath = 260 // MAX_PATH Traditional

func spi(action uint32, uiParam uintptr, pvParam unsafe.Pointer, flags uint32) error {
	r1, _, e1 := spiProc.Call(uintptr(action), uiParam, uintptr(pvParam), uintptr(flags))
	if r1 == 0 {
		if e1 != nil && !errors.Is(e1, windows.ERROR_SUCCESS) {
			return e1
		}
		return fmt.Errorf("SystemParametersInfoW failed (action=0x%x)", action)
	}
	return nil
}

func Get() (string, error) {
	buf := make([]uint16, maxPath+1)
	if err := spi(spiGetDeskWallpaper, uintptr(len(buf)), unsafe.Pointer(&buf[0]), 0); err != nil {
		return "", fmt.Errorf("get wallpaper: %w", err)
	}
	n := 0
	for n < len(buf) && buf[n] != 0 {
		n++
	}
	return string(utf16.Decode(buf[:n])), nil
}

func SetFromFile(path string) error {
	// NOTE: better to use ABSOLUTE path but it's not supported yet need looking in the future -- for now this works 👀

	p := filepath.Clean(path)

	if strings.ContainsAny(p, "%") {
		p = os.ExpandEnv(p)
	}
	if _, err := os.Stat(p); err != nil {
		return fmt.Errorf("stat: %w", err)
	}
	p16, err := windows.UTF16PtrFromString(p)
	if err != nil {
		return err
	}

	if err := spi(spiSetDeskWallpaper, 0, unsafe.Pointer(p16), spifUpdateINIFile|spifSendChange); err != nil {
		return fmt.Errorf("SPI_SETDESKWALLPAPER: %w", err)
	}
	return nil
}

type Mode int

const (
	Center Mode = iota
	Tile
	Stretch
	Fit
	Fill // Fill is the value not Crop
	Span
)

var modeMap = map[Mode]struct{ style, tile string }{
	Center:  {"0", "0"},
	Tile:    {"0", "1"},
	Stretch: {"2", "0"},
	Fit:     {"6", "0"},
	Fill:    {"10", "0"},
	Span:    {"22", "0"},
}

func SetMode(m Mode) error {
	vals, ok := modeMap[m]
	if !ok {
		return fmt.Errorf("invalid mode: %v", m)
	}

	key, _, err := registry.CreateKey(
		registry.CURRENT_USER,
		`Control Panel\Desktop`,
		registry.SET_VALUE,
	)
	if err != nil {
		return err
	}
	defer key.Close()

	if err := key.SetStringValue("TileWallpaper", vals.tile); err != nil {
		return err
	}
	if err := key.SetStringValue("WallpaperStyle", vals.style); err != nil {
		return err
	}

	cur, err := Get()
	if err != nil {
		return err
	}
	return SetFromFile(cur)
}

func Set(path string, m Mode) error {
	if err := SetFromFile(path); err != nil {
		return err
	}
	return SetMode(m)
}
