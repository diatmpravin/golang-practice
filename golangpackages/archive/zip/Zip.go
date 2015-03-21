package main

import (
	"archive/zip"
	"fmt"
)

func main() {
	read, err := zip.OpenReader("/home/elpsstu/gocode/src/golangpackages/archive/zip/test.zip")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(read)

	defer read.Close()
}
