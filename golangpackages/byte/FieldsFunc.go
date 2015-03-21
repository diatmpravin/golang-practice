package main

import (
	"bytes"
	"fmt"
)

func main() {
	strslice := []byte("a b c d e       f    \tg") // \t is tab space

	fmt.Println(string(strslice))

	strfield := bytes.FieldsFunc(strslice, func(divide rune) bool {
		return divide == ' ' // we divide at empty white space
	})

	for i, subslice := range strfield {
		fmt.Printf("Counter %d : %s \n", i, string(subslice))
	}

	utf8slice := []byte("大世界和小世界")

	utf8field := bytes.FieldsFunc(utf8slice, func(divide rune) bool {
		return divide == '和'
	})

	for i, utf8slice := range utf8field {
		fmt.Printf("Counter %d : %s \n", i, string(utf8slice))
	}

}
