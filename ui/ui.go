package ui

import (
	"fmt"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/0xdevar/waraqah"
	"github.com/0xdevar/waraqah/logic"
	"github.com/0xdevar/waraqah/repos"
)

func Start() {
	a := app.New()
	w := a.NewWindow("Hello World")

	tmp, _ := os.MkdirTemp("", "waraqah")
	log.Println(tmp)

	repo := repos.NewGitRepo("0xdevar", "0xwaraqat", "main", tmp)
	waraqahInstance, err := logic.RetrieveWallpapers(repo)

	// pages := waraqahInstance.Length() / 2

	var wallpapers []waraqah.WallpaperCollection

	if err != nil {
		log.Panicln("error", err)
		return
	}

	label := widget.NewLabel("")
	v := container.NewVBox()

	seek := func() {
		v.RemoveAll()
		var out []string
		for _, wallpaper := range wallpapers {
			out = append(out, fmt.Sprintf("image %s [images %d]", wallpaper.Name, len(wallpaper.Images)))
			repo.DownloadWallpaper(wallpaper)
			for _, img := range wallpaper.Images {
				fyne.Do(func() {
					image := canvas.NewImageFromFile(img.Path)

					image.FillMode = canvas.ImageFillContain
					image.SetMinSize(fyne.NewSize(100, 100))

					v.Add(image)
				})
			}
		}

		label.SetText(strings.Join(out, "\n"))
	}

	prevButton := widget.NewButton("Prev", func() {
		// if i < pages {
		// 	i = 0
		// }
		wallpapers = waraqahInstance.Prev(2)
		seek()
	})

	nextButton := widget.NewButton("Next", func() {
		// if i > pages {
		// 	i = pages
		// }

		wallpapers = waraqahInstance.Next(2)

		seek()
	})

	vs := container.NewVScroll(
		v,
	)

	vs.SetMinSize(fyne.NewSize(1000, 1000))

	w.SetContent(
		container.NewVBox(
			label,
			prevButton,
			nextButton,
			vs,
		))

	w.ShowAndRun()
}
