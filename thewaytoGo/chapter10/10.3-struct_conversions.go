// 10.3-struct_conversions.go
package main

import "fmt"

type number struct {
	f float32
}
type nr number // alias type

func main() {

	a := number{5.0}
	b := nr{5.0}

	//float32 in assignment
	//type float32
	//number in assignment

	var c = number(b)
	fmt.Println(a, b, c)
}
