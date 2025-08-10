package waraqah

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DownloadDir string `json:"download_dir"`
}

func LoadConfig() *Config {
	homedir, err := os.UserHomeDir()

	base := func(path string) string {
		return fmt.Sprintf("%s/%s", homedir, path)
	}

	if err != nil {
		return &Config{
			DownloadDir: base("waraqat"),
		}
	}

	configPath := base(".waraqah")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &Config{
			DownloadDir: base("waraqat"),
		}
	}

	content, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var c Config

	err = json.Unmarshal(content, &c)

	if err != nil {
		panic(err)
	}

	return &c
}
