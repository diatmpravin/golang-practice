package main

import (
	"fmt"
)

func checkPrime(n int) string {
	isPrimt := true
	if n == 1 || n == 2 || n == 3 {
		return "Prime"
	}

	for j := 2; j <= n/2+1; j++ {
		if n%j == 0 {
			isPrimt = false
			return "Not prime"
		}
	}
	if isPrimt {
		return "Prime"
	}
	return "Prime"
}

func main() {
	var t, n int

	fmt.Scan(&t)

	if t >= 1 && t <= 20 {
		for i := 0; i < t; i++ {
			fmt.Scan(&n)
			fmt.Println(checkPrime(n))
		}
	}
}
