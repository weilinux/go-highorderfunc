package main

import (
	"os"
	"path/filepath"
)
import "fmt"

func callback(path string, dir os.DirEntry, dirErr error) (err error) {
	info, _ := dir.Info()
	fmt.Printf("Path %v, Size: %v\n", path, info.Size())
	return
}


func main() {
	path, err := os.Getwd()
	if err == nil {
		err = filepath.WalkDir(path, callback)

	} else {
		fmt.Printf("Error %v", err.Error())
	}
	fmt.Println("main done")

}
