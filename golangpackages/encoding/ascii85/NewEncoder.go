package main

import (
	"bytes"
	"encoding/ascii85"
	"fmt"
)

func main() {

	var bytesbuf bytes.Buffer

	encoderstream := ascii85.NewEncoder(&bytesbuf)
	encoderstream.Write([]byte("abc123"))
	encoderstream.Close()
	fmt.Println(encoderstream)
}
