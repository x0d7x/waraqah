package ui

import (
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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

	repo := repos.NewGitRepo("0xdevar", "0xwaraqat", "main", tmp)
	waraqahInstance := logic.NewWaraqah(repo)

	maxPages := 1

	ch := make(chan []waraqah.WallpaperCollection, maxPages)
	index := make(chan int)

	label := widget.NewLabel("")

	prevButton := widget.NewButton("Prev", func() {
		index <- 0
	})

	nextButton := widget.NewButton("Next", func() {
		index <- 1
	})

	go func() {
		for item := range ch {
			for _, wallpaper := range item {
				fyne.Do(func() {
					label.SetText(fmt.Sprintf("image %s [images %d]", wallpaper.Name, len(wallpaper.Images)))
				})
			}
		}
	}()

	v := container.NewVBox(
		label,
		prevButton,
		nextButton,
	)

	go func() {
		err := waraqahInstance.Load(ch, index, maxPages)

		if err != nil {
			log.Panicln("error", err)
			return
		}
	}()

	w.SetContent(v)
	w.ShowAndRun()

}
