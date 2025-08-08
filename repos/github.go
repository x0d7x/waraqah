package repos

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GithubRepo struct {
	owner string
	repo  string
}

type githubFile struct {
	Name string `json:"name"`
}

func CreateRepo(owner, repo string) *GithubRepo {
	return &GithubRepo{owner, repo}
}

func (s *GithubRepo) getAllFiles(path *string) []string {
	var finalPath string
	if path == nil {
		finalPath = ""
	} else {
		finalPath = fmt.Sprintf("%s/", *path)
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/trees/master%s", s.owner, s.repo, finalPath)
	resp, err := http.Get(url)

	log.Println(resp.Status)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()

	var files []githubFile

	err = json.NewDecoder(resp.Body).Decode(&files)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for i, file := range files {
		fmt.Println(i+1, file.Name)
	}

	return nil
}

func (s *GithubRepo) GetAllFiles() []string {
	return s.getAllFiles(nil)
}

func (s *GithubRepo) GetAllFilesInPath(path string) []string {
	return s.getAllFiles(&path)
}
