package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "pravin"

	for i := 0; i < len(str); i++ {
		fmt.Println(reflect.TypeOf(str[i]))
	}

	fmt.Println("##############################")

	for i, v := range str {
		fmt.Println(reflect.TypeOf(v), i)
	}
}
