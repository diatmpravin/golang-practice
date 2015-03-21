// 7.1-for_arrays.go
package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("The value at index %d is: %d\n", i, arr1[i])
	}

}
