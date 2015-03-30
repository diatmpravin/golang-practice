package main

import (
	"fmt"
)

type HelloString string

func main() {
	const hello = "Hello, 世界"
	var s string
	s = hello

	acceptString(hello)
	acceptHelloString(hello)
	fmt.Println(s)

}

func acceptString(str string)          {}
func acceptHelloString(hs HelloString) {}