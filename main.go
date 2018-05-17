package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

var GOPATH = os.Getenv("GOPATH")

func main() {
	projectName := os.Args[1]
	user, _ := user.Current()
	projectDir := filepath.Join(GOPATH, "src/github.com/", user.Name, projectName)
	createProject(projectDir)
}

func createProject(dir string) {
	if _, err := os.Stat(dir); err == nil {
		fmt.Println("\"", dir, "\" is already exist")
		os.Exit(0)
	}
	if err := os.Mkdir(dir, 0775); err != nil {
		panic(err)
	}
	if err := os.Chdir(dir); err != nil {
		panic(err)
	}
	if err := exec.Command("git", "init").Run(); err != nil {
		panic(err)
	}
	fmt.Println("created", dir)
}
