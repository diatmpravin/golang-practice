package main

import (
	"fmt"
)

func main() {
	str := "hello-world"

	// gives only the bytes
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	// gives the Unicode characters
	for ix, ch := range str {
		fmt.Println(ix , " :" , ch)
	}
}
