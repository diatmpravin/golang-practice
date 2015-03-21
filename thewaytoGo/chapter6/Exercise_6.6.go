// Exercise_6.6.go
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 30; i++ {
		if i != 0 {
			factorial(i)
		}
	}
}

func factorial(n int) {
	fmt.Printf("Factorial of %d is: %d\n", n, n*(n-1))
}
