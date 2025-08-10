package logic

import (
	"fmt"

	"github.com/0xdevar/waraqah/repos"
)

type Waraqah struct {
	repo *repos.Git
}

func NewWaraqah(repo *repos.Git) Waraqah {
	return Waraqah{repo}
}

func (s *Waraqah) Load() {
	out, _ := s.repo.GetWallpapers()

	for i, f := range out {
		fmt.Printf("[%d] %s\n", i+1, f.Name)
	}

	var choice int
	fmt.Print("Select image number to download: ")
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(out) {
		fmt.Println("Invalid selection")
		return
	}

	selected := out[choice-1]
	if err := s.repo.DownloadWallpaper(selected); err != nil {
		fmt.Println("Download failed:", err)
	} else {
		fmt.Println("Downloaded:", selected.Name)
	}
}
