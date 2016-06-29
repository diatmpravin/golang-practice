package main

import (
	"fmt"
)

func sum(n int) (sum uint64) {
	var temp uint64

	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		sum += temp
	}
	return
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(sum(n))
}
