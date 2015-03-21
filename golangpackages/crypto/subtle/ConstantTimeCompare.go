package main

import (
	"crypto/subtle"
	"fmt"
)

func main() {
	X := []byte("Hello")
	Y := []byte("World")
	Z := []byte("Hello")

	n := subtle.ConstantTimeCompare(X, Y) // X != Y, n = 0
	fmt.Printf("n : %d\n", n)

	nn := subtle.ConstantTimeCompare(X, Z) // X == Z, nn = 1

	fmt.Printf("nn : %d\n", nn)
}
