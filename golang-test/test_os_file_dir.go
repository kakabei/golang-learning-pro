package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

func makeDir() {
	err := os.Mkdir("a", os.ModePerm)
	if err != nil {
		fmt.Printf("f dir err : %v\n", err)
	}
}

func checkFileName() {
	filename := "device/sdk/CMakeLists.acf"

	filesuffix := path.Ext(filename)
	pathDir := filepath.Dir(filename)
	fmt.Println("file filesuffix:", filesuffix)

	fmt.Println("file pathDir:", pathDir)
}

func main() {
	checkFileName()
}
