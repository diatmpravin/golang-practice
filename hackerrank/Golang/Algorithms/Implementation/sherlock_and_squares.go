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

		// b_sqrt := int64(math.Sqrt(float64(b)))
		// a_sqrt := int64(math.Sqrt(float64(a)))

		// count := b_sqrt - a_sqrt
		// if math.Sqrt(float64(a))-float64(a_sqrt) == 0 {
		// 	count++
		// }
		// fmt.Println(count)

		var count, i int64
		for i = a; i <= b; i++ {
			if math.Sqrt(float64(i)) - float64(int64(math.Sqrt(float64(i)))) == 0 {
				count++
			}
		}

		// if a == 1 {
		// 	count++
		// }

		fmt.Println(count)
	}
}
