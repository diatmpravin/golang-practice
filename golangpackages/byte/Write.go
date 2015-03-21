package main

import (
	"bytes"
	"fmt"
)

func main() {
	buff := bytes.NewBuffer(nil) // create empty buffer

	n, err := buff.Write([]byte("abc"))

	fmt.Printf("%s %d %v \n", string(buff.Bytes()), n, err)
}
