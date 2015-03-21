// Exercise_6.3.go

package main

import (
	"fmt"
)

func main() {
	myInt(1, 2, 3, 4, 5, 6, 7, 8)
}

func myInt(arg []int) {
	for v, _ := range arg {
		fmt.Println("Each Int is: ", v)
	}
}
