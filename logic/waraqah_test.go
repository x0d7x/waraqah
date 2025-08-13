package logic

import (
	"reflect"
	"testing"

	"github.com/0xdevar/waraqah"
)

var ws = []waraqah.WallpaperCollection{
	{Name: "1", Images: []waraqah.Wallpaper{{Path: "a"}}},
	{Name: "2", Images: []waraqah.Wallpaper{{Path: "b"}}},
	{Name: "3", Images: []waraqah.Wallpaper{{Path: "c"}}},
	{Name: "4", Images: []waraqah.Wallpaper{{Path: "d"}}},
	{Name: "5", Images: []waraqah.Wallpaper{{Path: "e"}}},
	{Name: "6", Images: []waraqah.Wallpaper{{Path: "f"}}},
	{Name: "7", Images: []waraqah.Wallpaper{{Path: "g"}}},
	{Name: "8", Images: []waraqah.Wallpaper{{Path: "h"}}},
	{Name: "9", Images: []waraqah.Wallpaper{{Path: "i"}}},
	{Name: "10", Images: []waraqah.Wallpaper{{Path: "j"}}},
}

func TestGetWallpapers(t *testing.T) {

	{
		w := Waraqah{wallpapers: ws, cursor: 0, chunkCount: 2}
		items := w.GetWallpapers(0)
		if reflect.DeepEqual(items, ws[0:2]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[0], items[0])
			return
		}

		items = w.GetWallpapers(2)
		if reflect.DeepEqual(items, ws[2:4]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[0], items[0])
			return
		}

		items = w.GetWallpapers(3)
		if reflect.DeepEqual(items, ws[3:5]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[0], items[0])
			return
		}

		items = w.GetWallpapers(4)
		if reflect.DeepEqual(items, ws[4:6]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[0], items[0])
			return
		}

		items = w.GetWallpapers(5)
		if reflect.DeepEqual(items, ws[5:7]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[0], items[0])
			return
		}

		items = w.GetWallpapers(11)
		if len(items) > 0 {
			t.Errorf("expected length > 0, got %v", len(items))
			return
		}

		items = w.GetWallpapers(10)
		if len(items) > 0 {
			t.Errorf("expected length > 0, got %v", len(items))
			return
		}
	}
}

func TestNext(t *testing.T) {

	{
		w := Waraqah{wallpapers: ws, cursor: 0, chunkCount: 2}

		items := w.GetWallpapers(0)
		if reflect.DeepEqual(items, ws[0:2]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[0], items[0])
			return
		}

		items = w.Next()
		if reflect.DeepEqual(items, ws[1:3]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[2:4], items)
			return
		}

		items = w.Next()
		if reflect.DeepEqual(items, ws[2:4]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[4:6], items)
			return
		}

		items = w.Next()
		if reflect.DeepEqual(items, ws[3:5]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[6:8], items)
			return
		}

		items = w.Next()
		if reflect.DeepEqual(items, ws[4:6]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[8:10], items)
			return
		}

		w.Next()
		w.Next()
		w.Next()
		w.Next()
		w.Next()
		items = w.Next()
		if len(items) != 0 {
			t.Errorf("expected length > 0, got %v", len(items))
			return
		}
	}
}

func TestPrev(t *testing.T) {

	{
		w := Waraqah{wallpapers: ws, cursor: 0, chunkCount: 2}

		items := w.GetWallpapers(5)
		if reflect.DeepEqual(items, ws[5:7]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[8:10], items)
			return
		}

		items = w.Prev()
		if reflect.DeepEqual(items, ws[4:6]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[6:8], items)
			return
		}

		items = w.Prev()
		if reflect.DeepEqual(items, ws[3:5]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[4:6], items)
			return
		}

		items = w.Prev()
		if reflect.DeepEqual(items, ws[2:4]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[2:4], items)
			return
		}

		items = w.Prev()
		if reflect.DeepEqual(items, ws[1:3]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[0:2], items)
			return
		}

		w.Prev()
		items = w.Prev()
		if len(items) != 0 {
			t.Errorf("expected length > 0, got %v", len(items))
			return
		}

		items = w.Prev()
		if len(items) != 0 {
			t.Errorf("expected length > 0, got %v", len(items))
			return
		}
	}
}

func TestNextPrev(t *testing.T) {
	{
		w := Waraqah{wallpapers: ws, cursor: 0, chunkCount: 4}

		items := w.GetWallpapers(1)
		if reflect.DeepEqual(items, ws[1:5]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[0], items[0])
		}

		items = w.Next()
		if reflect.DeepEqual(items, ws[2:6]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[2:4], items)
		}

		items = w.Prev()
		if reflect.DeepEqual(items, ws[1:5]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[0:2], items)
		}

		items = w.Next()
		if reflect.DeepEqual(items, ws[2:6]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[2:4], items)
		}

		items = w.Prev()
		if reflect.DeepEqual(items, ws[1:5]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[0:4], items)
		}

		items = w.Next()
		if reflect.DeepEqual(items, ws[2:6]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[2:4], items)
		}

		items = w.Next()
		if reflect.DeepEqual(items, ws[3:7]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[4:6], items)
		}

		items = w.Next()
		if reflect.DeepEqual(items, ws[4:8]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[6:8], items)
		}

		items = w.Next()
		if reflect.DeepEqual(items, ws[5:9]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[8:10], items)
		}

		for range 4 {
			w.Next()
		}

		items = w.Next()
		if len(items) != 0 {
			t.Errorf("expected length > 0, got %v", len(items))
		}

		items = w.Prev()
		if reflect.DeepEqual(items, ws[9:10]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[8:10], items)
		}

		items = w.Prev()
		if reflect.DeepEqual(items, ws[8:10]) != true {
			t.Errorf("expected %v, got %v", w.wallpapers[2:8], items)
		}
	}
}
