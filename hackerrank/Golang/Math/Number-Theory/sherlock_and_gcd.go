package main

import "fmt"

func GCD(a, b int) int {
	for b != 0 {
		r := a % b
		a, b = b, r
	}
	return a
}

func main() {
	var t, n int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		var a, b int

		fmt.Scan(&a)
		for j := 1; j < n; j++ {
			fmt.Scan(&b)
			a = GCD(a, b)
		}
		if a == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
