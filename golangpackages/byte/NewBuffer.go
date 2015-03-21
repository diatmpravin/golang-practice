package main

import (
	"bytes"
	"fmt"
)

func main() {
	newbuf := bytes.NewBuffer([]byte("abcdefg"))

	fmt.Println(newbuf)

	var buff [7]byte

	newbuf.Read(buff[0:5])

	fmt.Println(string(buff[:]))
}
