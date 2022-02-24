package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

// IsLink checks if the file is a link
func IsLink(path string) (string, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		return "", err
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Mode()&fs.ModeSymlink == fs.ModeSymlink)
	return fileInfo.Mode().String(), nil
}

func main() {
	res, err := IsLink("/home/workspace/shiyuelc/ltl")
	if err != nil {
		return
	}
	fmt.Println(res)
	path, _ := filepath.EvalSymlinks("/home/workspace/shiyuelc/ltl")
	fmt.Println(path)
}
