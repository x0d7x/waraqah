package logic

import (
	"github.com/0xdevar/waraqah"
	"github.com/0xdevar/waraqah/repos"
)

type Waraqah struct {
	repo *repos.Git
}

func NewWaraqah(repo *repos.Git) Waraqah {
	return Waraqah{repo}
}

func (s *Waraqah) Load(
	producer chan<- []waraqah.WallpaperCollection,
	index <-chan int,
	maxChunk int,
) error {
	wallpapers, err := s.repo.GetWallpapers()

	if err != nil {
		return err
	}

	var pageSize = maxChunk % len(wallpapers)

	for {
		i, ok := <-index

		if !ok {
			return nil
		}

		if i >= len(wallpapers) || i < 0 {
			i = 0
		}

		println(i)

		wallpapers := wallpapers[i*pageSize : ((i + 1) * pageSize)]

		producer <- wallpapers
	}
}
