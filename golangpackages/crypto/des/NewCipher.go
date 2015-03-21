package main

import (
	"crypto/des"
	"fmt"
)

func main() {

	key := "12345678" // 8 bytes! this is the DES block size in bytes

	block, err := des.NewCipher([]byte(key))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d bytes NewCipher key with block size of %d bytes\n", len(key), block.BlockSize)
}
