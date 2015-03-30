package main

import (
	"math"
	"fmt"
)

func main() {
	// Valid
	var a int64 = math.MaxInt64
	b := interface{}(int64(math.MaxInt64))

	// Not Valid
	c := math.MaxInt64
	d := interface{}(math.MaxInt64)

	fmt.Println(a,b)
	fmt.Println(c,d)
}
