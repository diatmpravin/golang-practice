package main

import (
	"bytes"
	"fmt"
)

func main() {
	buff := bytes.NewBuffer([]byte("abcdefg"))

	fmt.Println(buff.String())

	var buf *bytes.Buffer

	fmt.Println(buf.String())
}
