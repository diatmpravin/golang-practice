package main

import (
	"crypto/subtle"
	"fmt"
)

func main() {
	X := uint8(1)
	Y := uint8(2)

	n := subtle.ConstantTimeByteEq(X, Y) // X != Y, n = 0
	fmt.Printf("n : %d\n", n)

	X = X + 1

	nn := subtle.ConstantTimeByteEq(X, Y) // X == Y, nn = 1

	fmt.Printf("nn : %d\n", nn)
}
