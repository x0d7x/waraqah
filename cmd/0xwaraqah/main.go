package main

import "github.com/0xdevar/waraqah/ui"

func main() {
	ui.Start()
	// config := waraqah.LoadConfig()
	// repo := repos.NewGitRepo("0xdevar", "0xwaraqat", "main", config.DownloadDir)
	// w := logic.NewWaraqah(repo)
	//
	// ch := make(chan []waraqah.WallpaperCollection)
	//
	// index := make(chan int)
	//
	// go w.Load(ch, index, 1)
	//
	// // imagnie this from a ui where user can click next
	// // and it's increase the requests so it fetch next page
	// go func() {
	// 	for item := range ch {
	// 		for _, wallpaper := range item {
	// 			println("file", wallpaper.Name)
	// 		}
	// 	}
	// 	println("done")
	// }()
	//
	// go func() {
	// 	index <- 0
	// 	time.Sleep(time.Second * 3)
	// 	index <- 1
	// 	time.Sleep(time.Second * 1)
	// 	index <- 0
	// 	time.Sleep(time.Second * 1)
	// 	close(index)
	// 	close(ch)
	// 	// time.Sleep(time.Second * 1)
	// 	// index <- 4
	// }()
	//
	// time.Sleep(time.Second * 20)
}
