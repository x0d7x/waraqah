package main

import (
	"fmt"
	"log"
	"os"

	"github.com/0xdevar/waraqah"
	"github.com/0xdevar/waraqah/logic"
	"github.com/0xdevar/waraqah/repos"
)

func main() {
	config := waraqah.LoadConfig()
	fmt.Println(config)
	dir, _ := os.MkdirTemp("", "waraqah")
	log.Println(dir)
	repo := repos.NewGitRepo("0xdevar", "0xwaraqs", "main", config.DownloadDir)
	w := logic.NewWaraqah(repo)
	w.Load()
}
