// Exercise_6.2_1.go

package main

import (
	"fmt"
	"math"
)

func main() {
	result, err := MySqrt(-9)
	if len(err) != 0 {
		fmt.Println("Error:", err)
	}
	fmt.Println("MySqrt Root: ", result)
}

func MySqrt(x float64) (float64, string) {
	var err string
	if x < 0 {
		err = "Value is negative"
	}
	result := math.Sqrt(x)
	return result, err
}
