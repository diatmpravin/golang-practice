package main

import (
	"bytes"
	"fmt"
)

func main() {

	buff := bytes.NewBuffer([]byte("abcdefg"))

	buff.Truncate(2) // keep first 2 bytes and discard the rest

	fmt.Println(buff.String())

}
