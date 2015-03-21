package main

import (
	"crypto/aes"
	"fmt"
)

func main() {

	//The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
	key := "opensesame123456" // 16 bytes!

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(block)

	fmt.Printf("%d bytes NewCipher key with block size of %d bytes\n", len(key), block.BlockSize)
}
