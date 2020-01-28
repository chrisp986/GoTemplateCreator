package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFolderExist(path string) bool {
	var folderExists = false
	folderInfo, _ := os.Stat(path)
	if folderInfo != nil {
		log.Println("Folder already exists.")
		folderExists = true
	}
	return folderExists
}

func getProjectName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Project Name:")
	projectName, _ := reader.ReadString('\n')
	projectName = strings.Replace(projectName, "\r\n", "", -1)
	return projectName
}

func createFolders(projectName string) {
	pathBase := filepath.Join("C:/BASE8/" + projectName)
	var strEndings = []string{".c8", ".cwy", ".p8k", ".sys"}
	for _, ending := range strEndings {
		if checkFolderExist(pathBase+ending) == false {
			err := os.MkdirAll(pathBase+ending, os.ModePerm)
			check(err)
			fmt.Println("->", pathBase+ending, " created..")
		}
	}
}

func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func main() {
	fmt.Println("///  Go Template Creator v0.1  ///")
	projectName := getProjectName()
	createFolders(projectName)
}
