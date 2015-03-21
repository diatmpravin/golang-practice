// Example_4.go
package main

import (
	"fmt"
	"math"
)

func main() {
	t, OK := mySqrt(9)

	if OK {
		fmt.Println("Square of 9 is: ", t)
	}
}

func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		return
	}
	return math.Sqrt(f), true
}
