package main

import (
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	h256 := sha256.New()
	io.WriteString(h256, "Hello money money!")
	fmt.Printf("Hash256 : %x\n", h256.Sum(nil))

	h224 := sha256.New224()
	io.WriteString(h224, "Hello money money!")
	fmt.Printf("Hash224 : %x\n", h224.Sum(nil))

	// const Size224 = 28 bytes
	data224 := []byte("Hello World!")
	fmt.Printf("SHA224 checksum : %x\n", sha256.Sum224(data224))

	data := []byte("Hello World!") // const Size = 32 bytes
	fmt.Printf("SHA256 checksum : %x\n", sha256.Sum256(data))
}
