package main

import (
	"bufio"
	"fmt"
	"os"
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
	pathBase := ("BASE8" + projectName + "c8")
	// pathBase := filepath.Join("BASE8/", projectName, ".c8")

	fmt.Println(pathBase)
	// pathC8 := filepath.Join(pathBase + projectName + ".c8")
	// pathCwy := filepath.Join("BASE8/" + projectName + ".cwy")
	// pathSys := filepath.Join("BASE8/" + projectName + ".sys")
	// var strPaths = []string{pathC8, pathCwy, pathSys}

	// for _, path := range strPaths {
	// 	err := os.MkdirAll(path, os.ModePerm)
	// 	check(err)
	// 	fmt.Println(path, " created..")
	// }
}

func main() {
	fmt.Println("///  Go Template Creator v0.1  ///")
	projectName := getProjectName()
	createFolders(projectName)
}
