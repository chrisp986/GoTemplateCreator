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

const basePath = "C:/BASE8/"
const dwgFilePath = "temp_files\\Blank_100_100.dwg"
const cnfFilePath = "temp_files\\temp.cnf"

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

func getInput() (output string) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.Replace(input, "\r\n", "", -1)
}

func getProjectName() string {
	fmt.Print("Enter Project Name:")
	userInput := getInput()
	return userInput
}

func createFolders(projectName string) {
	fullPath := filepath.Join(basePath + projectName)
	var strEndings = []string{".c8", ".cwy", ".p8k", ".sys"}
	for _, ending := range strEndings {
		if checkFolderExist(fullPath+ending) == false {
			err := os.MkdirAll(fullPath+ending, os.ModePerm)
			check(err)
			fmt.Println("->", fullPath+ending, " created..")
		}
	}
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func main() {
	fmt.Println("///  Go Template Creator v0.1  ///")
	projectName := getProjectName()
	createFolders(projectName)
	copy(dwgFilePath, basePath+projectName+".sys/Blank_100_100.dwg")
	fmt.Println("dwg file copied...")
	copy(cnfFilePath, basePath+projectName+".p8k/"+projectName+".cnf")
	fmt.Println("cnf file copied...")
}
