package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	str := []byte("input string to be read into new buffer")
	readbuffer := bytes.NewBuffer(str)

	write := bufio.NewWriter(os.Stdout)

	readwrite := bufio.NewReadWriter(bufio.NewReader(readbuffer), bufio.NewWriter(write))
	fmt.Printf("%#v",readwrite)
}
