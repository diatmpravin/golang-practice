package main

import (
	"container/ring"
	"fmt"
)

func main() {

	r := ring.New(10)

	// populate our ring
	for i := 0; i < 10; i++ {
		r.Value = i
		r = r.Next()
	}

	r2 := ring.New(10)

	// populate our ring
	for t := 0; t < 10; t++ {
		r2.Value = t
		r2 = r2.Next()
	}

	f := func(v interface{}) {
		fmt.Printf("%d ", v)
	}

	// Values in r before link
	fmt.Println("Values in r BEFORE link")
	r.Do(f)
	fmt.Println()

	r.Link(r2)

	// Values in r AFTER link
	fmt.Println("Values in r AFTER link")
	r.Do(f)
	fmt.Println()

}
