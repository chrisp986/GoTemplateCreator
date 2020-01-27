package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getProjectName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Project Name:")
	projectName, _ := reader.ReadString('\n')
	return projectName
}

func createFolders(projectName string) {
	pathC8 := filepath.Join("Base8/" + projectName + ".c8")
	pathCwy := filepath.Join("Base8/" + projectName + ".cwy")
	pathSys := filepath.Join("Base8/" + projectName + ".sys")
	paths := list.New()
	paths.PushFront(pathC8)
	paths.PushFront(pathCwy)
	paths.PushFront(pathSys)

	for index := 0; index < paths.Len(); index++ {
		err := os.MkdirAll(pathC8, os.ModePerm)

	}
	err := os.MkdirAll(pathC8, os.ModePerm)
	err := os.MkdirAll(pathCwy, os.ModePerm)
	err := os.MkdirAll(pathSys, os.ModePerm)
	check(err)
}

func main() {
	fmt.Println("///  Go Template Creator v0.1  ///")
	projectName := getProjectName()
	createFolders(projectName)
}
