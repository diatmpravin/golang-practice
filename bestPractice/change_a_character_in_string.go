package main

import (
	"fmt"
)

func main() {
	str := "hello"
	c := []byte(str)
	fmt.Println(c)
	c[0] = 'c'
	fmt.Println(c)
	s2 := string(c)
	fmt.Println("--->", s2)
}
