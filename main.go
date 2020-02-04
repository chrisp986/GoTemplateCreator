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

//createFolder() takes the projectName as input and adds the strings to the foldernames to create a full path
//the new folders are then created in the C:\BASE8\ folder, since this is fixed this is a constant variable
//Example: Project.c8, Project.cwy, Project.p8k, Project.sys
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

//checkFileExist() is called with the path it needs to check and then returns a bool value based on the os.FileInfo value
func checkFileExist(path string) bool {
	var fileExist = false
	_, err := os.Stat(path)
	if err == nil {
		log.Println(path, "already exists.")
		fileExist = true
	}
	return fileExist
}

//getInput() takes the input from the user and returns this as string
func getInput() (output string) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.Replace(input, "\r\n", "", -1)
}

//getProjectName() asks for the projectName
func getProjectName() string {
	fmt.Print("Enter Project Name:")
	userInput := getInput()
	return userInput
}

//copy() checks if the file/folder exists and base on that copies the file to the folders in the C:\BASE8\ folder
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

//main() the user has to enter the project name which will be used to name and then create the folders
//after that the blnk dwg file and the renamed cnf file are copied in the new folders
func main() {
	fmt.Println("///  Go Template Creator v0.1  ///")
	projectName := getProjectName()
	createFolders(projectName)
	copy(dwgFilePath, basePath+projectName+".sys/Blank_100_100.dwg")
	copy(cnfFilePath, basePath+projectName+".p8k/"+projectName+".cnf")
	time.Sleep(5 * time.Second)
}
