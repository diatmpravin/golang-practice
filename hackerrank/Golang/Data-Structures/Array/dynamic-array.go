package main

import "fmt"

func main() {
	var n int
	var q int

	var a [][]int

	var lastAnswer = 0

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var ea []int

		a = append(a, ea)
	}

	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var t int
		var x int
		var y int

		fmt.Scan(&t)
		fmt.Scan(&x)
		fmt.Scan(&y)

		if t == 1 {
			a[(x ^ lastAnswer) % n] = append(a[(x ^ lastAnswer) % n], y)
		} else {
			res := a[(x ^ lastAnswer) % n][y % len(a[(x ^ lastAnswer) % n])]

			lastAnswer = res

			fmt.Println(res)
		}
	}
}
