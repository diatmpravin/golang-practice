// 6.2-simple_function.go

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Multiply of 2*3*4= %d \n", multiPly3Num(2, 3, 4))
}

func multiPly3Num(a int, b int, c int) int {
	return a * b * c
}
