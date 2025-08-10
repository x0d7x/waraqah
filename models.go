package waraqah

type WallpaperMeta struct {
	Resolution [2]int
	Size       int
}

type Wallpaper struct {
	WallpaperMeta
	Path string
}

type WallpaperCollection struct {
	Name   string
	Images []Wallpaper
}
