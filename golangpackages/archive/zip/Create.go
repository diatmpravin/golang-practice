package main

import (
	"archive/zip"
	"bytes"
	"fmt"
)

func main() {
	buffer := new(bytes.Buffer)

	fileInfo := file.Stat()
	header := zip.FileInfoHeader(fileInfo)
	writer := zipWriter.CreateHeader(header)

	zipWriter := zip.NewWriter(buffer)
	writer, err := zipWriter.Create("newfile.zip")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(writer)
}
