package main

import (
	"bytes"
	"fmt"
)

func main() {
	newbuf := bytes.NewBufferString("abcdefg") // just a string will do

	var buff [7]byte

	newbuf.Read(buff[:])

	fmt.Println(string(buff[:]))
}
