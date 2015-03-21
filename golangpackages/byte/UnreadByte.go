package main

import (
	"bytes"
	"fmt"
)

func main() {

	buff := bytes.NewBuffer([]byte("abcdefg"))

	fmt.Println(string(buff.Next(3))) // read next 3 bytes

	fmt.Println(buff.String()) // starts from d

	buff.UnreadByte() // reset read position back to c (1 byte)
	fmt.Println(buff.String())
}
