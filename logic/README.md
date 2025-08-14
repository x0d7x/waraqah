# Waraqah OS-Specific Logic

This document explains the functionality of the OS-specific logic for Waraqah, focusing on the Windows implementation (`os_win.go`). The same principles apply to `os_linux.go` and `os_mac.go` for their respective operating systems.

## Windows Implementation (`os_win.go`)

The `os_win.go` file provides the necessary functions to interact with the Windows operating system to manage the desktop wallpaper.

### Functions

#### `Get() (string, error)`

This function retrieves the path of the current desktop wallpaper.

**Usage:**

```go
currentWallpaper, err := Get()
if err != nil {
    // Handle error
}
fmt.Println("Current wallpaper:", currentWallpaper)
```

#### `SetFromFile(path string) error`

This function sets the desktop wallpaper to the image specified by the `path`.

**Usage:**

```go
err := SetFromFile("C:\path\to\your\wallpaper.jpg")
if err != nil {
    // Handle error
}
```

#### `SetMode(m Mode) error`

This function sets the wallpaper display mode (e.g., Center, Tile, Stretch).

**Available Modes:**

*   `Center`
*   `Tile`
*   `Stretch`
*   `Fit`
*   `Fill`
*   `Span`

**Usage:**

```go
err := SetMode(Stretch)
if err != nil {
    // Handle error
}
```

#### `Set(path string, m Mode) error`

This function is a convenience wrapper that first sets the wallpaper from a file and then sets the display mode.

**Usage:**

```go
err := Set("C:\path\to\your\wallpaper.jpg", Fill)
if err != nil {
    // Handle error
}
```

## Linux and macOS Implementations

The `os_linux.go` and `os_mac.go` files provide equivalent functionality for Linux and macOS, respectively. They expose the same function signatures for getting and setting the wallpaper, allowing for cross-platform compatibility at a higher level in the Waraqah application.

