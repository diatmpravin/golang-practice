package main

import (
	"bytes"
	"fmt"
)

func main() {
	buff := bytes.NewBufferString("abcdefghi")

	var threebytes [3]byte

	buff.Read(threebytes[0:3])

	fmt.Printf("%s\n", string(threebytes[:]))

	// read the next 5 bytes

	fmt.Printf("%s\n", string(buff.Next(5)))

}
