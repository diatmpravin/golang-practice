package main

import (
	"crypto/subtle"
	"fmt"
)

func main() {
	x := int32(1)
	y := int32(2)
	z := int32(1)

	n := subtle.ConstantTimeEq(x, y)
	fmt.Printf("n : %d \n", n)

	nn := subtle.ConstantTimeEq(x, z)
	fmt.Printf("nn : %d \n", nn)

}
