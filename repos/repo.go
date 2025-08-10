package repos

import w "github.com/0xdevar/waraqah"

type WallpaperRepo interface {
	GetWallpapers() ([]w.WallpaperCollection, error)
	DownloadWallpaper(wallpaper w.WallpaperCollection) error
}
