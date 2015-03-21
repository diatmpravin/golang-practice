package main

import (
	"bytes"
	"fmt"
)

func main() {

	buff := bytes.NewBuffer([]byte("abcdefg"))

	fmt.Printf("%d \n", buff.Len())

	// remove all content from the buff buffer
	buff.Reset()

	fmt.Printf("%d \n", buff.Len())
}
