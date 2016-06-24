package main

import (
	"bytes"
	"fmt"
)

func main() {

	search_input := []byte("abcdefghidefdef")
    fmt.Println("---->", search_input)

    key_slice := []byte("def")
    fmt.Println("---->", key_slice)

	fmt.Println(index(search_input, key_slice))
}

// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.
func index(s, sep []byte) int {
	n := len(sep)
	if n == 0 {
		return 0
	}
	if n > len(s) {
		return -1
	}
	c := sep[0]
	if n == 1 {
		return bytes.IndexByte(s, c)
	}
	i := 0
	t := s[:len(s)-n+1]
	for i < len(t) {
		if t[i] != c {
			o := bytes.IndexByte(t[i:], c)
			if o < 0 {
				break
			}
			i += o
		}
		if bytes.Equal(s[i:i+n], sep) {
			return i
		}
		i++
	}
	return -1
}
