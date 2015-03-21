package main

import (
	"bytes"
	"fmt"
)

func main() {
	buff := bytes.NewBufferString("abcdefghi")

	fmt.Printf("%d\n", len(buff.Bytes()))

	fmt.Printf("%d\n", buff.Len())

}
