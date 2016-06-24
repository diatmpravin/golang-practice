package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	slice := []byte("abcdefg")
	slicebyteindex := indexRune(slice, 'd')

	fmt.Printf("d index num is %d inside the string %s\n", slicebyteindex, string(slice))

	utf8slice := []byte("大世界和小世界")

	utf8byteindex := indexRune(utf8slice, '小')
	fmt.Printf("小 index num is %d inside the string %s\n", utf8byteindex, string(utf8slice))

	utf8byteindex2 := indexRune(utf8slice, '界')
	fmt.Printf("界 index num is %d inside the string %s\n", utf8byteindex2, string(utf8slice))
}

// IndexRune interprets s as a sequence of UTF-8-encoded Unicode code points.
// It returns the byte index of the first occurrence in s of the given rune.
// It returns -1 if rune is not present in s.
func indexRune(s []byte, r rune) int {
	for i := 0; i < len(s); {
		r1, size := utf8.DecodeRune(s[i:])
		if r == r1 {
			return i
		}
		i += size
	}
	return -1
}
