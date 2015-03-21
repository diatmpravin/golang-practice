package main

import (
	"encoding/ascii85"
	"fmt"
)

func main() {
	str := []byte("Man is distinguished, not only by his reason, but by this singular passion from")
	fmt.Println(len(str))
	buffer := make([]byte, ascii85.MaxEncodedLen(len(str)))
	fmt.Println(len(buffer))
	encodedbytes := ascii85.Encode(buffer, str)

	fmt.Println("Bytes written : ", encodedbytes)
	fmt.Println(string(buffer))
}
