package main

import (
	"fmt"
	"math"
)

func main() {
	var t int

	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b int64
		fmt.Scan(&a, &b)

		b_sqrt := int64(math.Sqrt(float64(b)))
		a_sqrt := int64(math.Sqrt(float64(a)))

		count := b_sqrt - a_sqrt
		if math.Sqrt(float64(a))-float64(a_sqrt) == 0 {
			count++
		}
		fmt.Println(count)
	}
}
