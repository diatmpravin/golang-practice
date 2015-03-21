// 6.4-blank_identifier.go
package main

import (
	"fmt"
)

func main() {
	var i1 int
	var f1 float64
	i1, _, f1 = threeValue()
	fmt.Println(i1, f1)
}

func threeValue() (int, int, float64) {
	return 2, 3, 1.2
}
