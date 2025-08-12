package logic

import (
	"testing"

	"github.com/0xdevar/waraqah"
)

func TestSeek(t *testing.T) {
	check := func(cases ...[5]int) {
		w := &Waraqah{wallpapers: make([]waraqah.WallpaperCollection, 100), page: 0}
		for _, c := range cases {
			page, length, wantStart, wantEnd, wantPage := c[0], c[1], c[2], c[3], c[4]
			start, end := w._seek(page, length)
			if start != wantStart || end != wantEnd || w.page != wantPage {
				t.Errorf("seek(%d, %d) = (%d, %d, %d), want (%d, %d, %d)",
					page, length, start, end, w.page, wantStart, wantEnd, wantPage)
			}
		}
	}

	check(
		[5]int{1, 2, 0, 2, 1},
		[5]int{2, 2, 2, 4, 2},
	)

	check(
		[5]int{1, 50, 0, 50, 1},
		[5]int{2, 50, 50, 100, 2},
		[5]int{3, 50, -1, -1, 2},
	)

	check(
		[5]int{0, 10, -1, -1, 0},
	)
}
