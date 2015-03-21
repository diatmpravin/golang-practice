package main

import (
	"bytes"
	"fmt"
)

func main() {
	buff := bytes.NewBufferString("abcd*ef")
	rb, err := buff.ReadBytes('*')
	fmt.Printf("%v  %d  %s\n", err, rb, string(rb))
}
