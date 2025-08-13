package logic

import (
	"github.com/0xdevar/waraqah"
	"github.com/0xdevar/waraqah/repos"
)

type Waraqah struct {
	wallpapers []waraqah.WallpaperCollection
	cursor     int
	chunkCount int
}

func RetrieveWallpapers(repo *repos.Git, chunkCount int) (Waraqah, error) {
	wallpapers, err := repo.GetWallpapers()

	if err != nil {
		return Waraqah{}, err
	}

	waraqah := Waraqah{wallpapers, 0, chunkCount}

	return waraqah, nil
}

func (s *Waraqah) _seek(page, length int) (int, int) {
	if page < 0 || length < 0 {
		return -1, -1
	}

	pageIndex := page

	if pageIndex < 0 {
		return -1, -1
	}

	start := pageIndex
	end := start + length

	if start > len(s.wallpapers) {
		return -1, -1
	}

	if start > len(s.wallpapers) {
		start = len(s.wallpapers) - length
	}

	if end > len(s.wallpapers) {
		end = len(s.wallpapers)
	}

	s.cursor = page

	return start, end
}

func (s *Waraqah) Length() int {
	return len(s.wallpapers)
}

func (s *Waraqah) GetWallpapers(page int) []waraqah.WallpaperCollection {
	start, end := s._seek(page, s.chunkCount)

	if start == -1 {
		return []waraqah.WallpaperCollection{}
	}

	return s.wallpapers[start:end]
}

func (s *Waraqah) Next() []waraqah.WallpaperCollection {
	i := s.cursor + 1
	start, end := s._seek(i, s.chunkCount)

	if start == -1 {
		return []waraqah.WallpaperCollection{}
	}

	return s.wallpapers[start:end]
}

func (s *Waraqah) Prev() []waraqah.WallpaperCollection {
	i := s.cursor - 1
	start, end := s._seek(i, s.chunkCount)

	if start == -1 {
		return []waraqah.WallpaperCollection{}
	}

	return s.wallpapers[start:end]
}
