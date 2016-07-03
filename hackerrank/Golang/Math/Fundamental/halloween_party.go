package main

import "fmt"

func main() {
	var cases, c int

	fmt.Scan(&cases)

	for i := 0; i < cases; i++ {
		fmt.Scan(&c)

		v := c / 2
		h := c - v
		fmt.Println(v * h)
	}
}
