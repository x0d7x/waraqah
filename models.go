package waraqah

type WallpaperSubmitter struct {
	ID   string
	Name string
}

type WallpaperMeta struct {
	Resolution [2]int
	Size       int
	Tags       []string
}

type Wallpaper struct {
	Filename  string
	By        *WallpaperSubmitter
	CreatedAt int
	Meta      WallpaperMeta
}
