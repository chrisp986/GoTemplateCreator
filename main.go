// Go Template Creator v0.1
// This tool is used to create a template project for KM projects which are used for automated vehicles.
// Copyright by Christian
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
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

func createFolders(projectName string) {
	fullPath := filepath.Join(basePath + projectName)
	var strEndings = []string{".c8", ".cwy", ".p8k", ".sys"}
	for _, ending := range strEndings {
		if checkFileExist(fullPath+ending) == false {
			err := os.MkdirAll(fullPath+ending, os.ModePerm)
			check(err)
			fmt.Println("->", fullPath+ending, " created..")
		}
	}
}

func checkFileExist(path string) bool {
	var fileExist = false
	_, err := os.Stat(path)
	if err == nil {
		log.Println(path, "already exists.")
		fileExist = true
	}
	return fileExist
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

func copy(src, dst string) (int64, error) {
	_, err := os.Stat(src)
	check(err)
	source, err := os.Open(src)
	check(err)
	defer source.Close()
	_, err = os.Stat(dst)
	if err == nil {
		fmt.Println("Skipping because", dst, "already exists.")
		return 0, err
	}
	destination, err := os.Create(dst)
	check(err)
	defer destination.Close()
	fmt.Println("-> Moved file to: ", dst)
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func main() {
	fmt.Println("///  Go Template Creator v0.1  ///")
	projectName := getProjectName()
	createFolders(projectName)
	copy(dwgFilePath, basePath+projectName+".sys/Blank_100_100.dwg")
	copy(cnfFilePath, basePath+projectName+".p8k/"+projectName+".cnf")
	time.Sleep(5 * time.Second)
}
