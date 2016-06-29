package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	leftIndex := 0
	rightIndex := n - 1
	sum := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var a int
			fmt.Scanf("%d", &a)
			if j == leftIndex {
				sum += a
			}
			if j == rightIndex {
				sum -= a
			}
		}

		leftIndex++
		rightIndex--
	}

	fmt.Println(math.Abs(float64(sum)))
}
