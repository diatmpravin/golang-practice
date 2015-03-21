package main

import (
	"fmt"
)

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func main() {
	// It works fine here, as I wrap things with array.
	c := []byte{'A'}
	fmt.Println(unhex(c[0]))

	d := byte('A') //    **Error** invalid type for composite literal: byte
	fmt.Println(unhex(d))
}
