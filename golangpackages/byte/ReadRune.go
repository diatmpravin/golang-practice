package main

import (
	"bytes"
	"fmt"
)

func main() {

	buff := bytes.NewBuffer([]byte("是你吗？"))

	rune, size, err := buff.ReadRune()
	fmt.Println(string(rune), size, err)

	// read next rune
	rune, size, err = buff.ReadRune()
	fmt.Println(string(rune), size, err)

	rune, size, err = buff.ReadRune()
	fmt.Println(string(rune), size, err)
	
	rune, size, err = buff.ReadRune()
	fmt.Println(string(rune), size, err)
	
	rune, size, err = buff.ReadRune()
	fmt.Println(string(rune), size, err)
}
