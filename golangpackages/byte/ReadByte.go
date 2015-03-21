package main

import (
	"bytes"
	"fmt"
)

func main() {
	buff := bytes.NewBufferString("abcdef")

	rb, err := buff.ReadByte()

	fmt.Printf("%v  %d  %s\n", err, rb, string(rb))
}
