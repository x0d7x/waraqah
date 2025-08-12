package logic

import (
	"github.com/0xdevar/waraqah"
	"github.com/0xdevar/waraqah/repos"
)

type Waraqah struct {
	wallpapers []waraqah.WallpaperCollection
	page       int
}

func RetrieveWallpapers(repo *repos.Git) (Waraqah, error) {
	wallpapers, err := repo.GetWallpapers()

	if err != nil {
		return Waraqah{}, err
	}

	waraqah := Waraqah{wallpapers, 1}

	return waraqah, nil
}

func (s *Waraqah) _seek(page, length int) (int, int) {
	if page < 0 || length < 0 {
		return -1, -1
	}

	pageIndex := page - 1

	if pageIndex < 0 {
		return -1, -1
	}

	start := pageIndex * length
	end := start + length

	if start == len(s.wallpapers) {
		return -1, -1
	}

	if start >= len(s.wallpapers) {
		start = len(s.wallpapers) - length
	}

	if end > len(s.wallpapers) {
		end = len(s.wallpapers)
	}

	s.page = page

	return start, end
}

func (s *Waraqah) Length() int {
	return len(s.wallpapers)
}

func (s *Waraqah) GetWallpapers(page, length int) []waraqah.WallpaperCollection {
	start, end := s._seek(page, length)

	if start == -1 {
		return []waraqah.WallpaperCollection{}
	}

	return s.wallpapers[start:end]
}

func (s *Waraqah) Next(length int) []waraqah.WallpaperCollection {
	i := s.page + 1
	start, end := s._seek(i, length)
	println(i, s.page, start, end)

	if start == -1 {
		return []waraqah.WallpaperCollection{}
	}

	return s.wallpapers[start:end]
}

func (s *Waraqah) Prev(length int) []waraqah.WallpaperCollection {
	i := s.page - 1
	start, end := s._seek(i, length)

	if start == -1 {
		return []waraqah.WallpaperCollection{}
	}

	return s.wallpapers[start:end]
}
