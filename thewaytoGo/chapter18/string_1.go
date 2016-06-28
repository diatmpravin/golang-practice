package main

import (
	"fmt"
)

func main() {
	str := "hello"
	c := []byte(str)
	c[0] = 'c'
	str1 := string(c)
	fmt.Println(str1)
}
