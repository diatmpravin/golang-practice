package main

import (
	"bytes"
	"fmt"
)

func main() {
	buff := bytes.NewBufferString("abcdefghi")

	fmt.Printf("%d\n", len(buff.Bytes()))

	buff.Grow(3)            // the buffer before writing. Good practice to allocate space before writing.
	buff.WriteString("def") // incase you forgot to grow the buffer's capacity. WriteString will automatically allocate space before writing into the buffer

	fmt.Printf("%d\n", len(buff.Bytes()))
}
