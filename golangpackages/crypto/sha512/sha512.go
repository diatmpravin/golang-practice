package main

import (
	"crypto/sha512"
	"fmt"
	"io"
)

func main() {
	h512 := sha512.New()
	io.WriteString(h512, "Hello money money!")
	fmt.Printf("Hash512 : %x\n", h512.Sum(nil))

	h384 := sha512.New384()
	io.WriteString(h384, "Hello money money!")
	fmt.Printf("Hash384 : %x\n", h384.Sum(nil))

	data := []byte("Hello World!")
	fmt.Printf("SHA384 checksum : %x\n", sha512.Sum384(data))

	fmt.Printf("SHA512 checksum : %x\n", sha512.Sum512(data))
}
