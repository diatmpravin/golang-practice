// 6.5-minmax.go
package main

import (
	"fmt"
)

func main() {
	var min, max int
	min, max = minMax(23, 12)
	fmt.Printf("Minimum is: %d, Maximum is: %d\n", min, max)
}

func minMax(a, b int) (min int, max int) {
	if a > b {
		min = b
		max = a
	} else {
		min = a
		max = b
	}
	return
}
