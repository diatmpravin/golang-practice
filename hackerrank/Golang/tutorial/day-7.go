package main

import (
	"fmt"
)

func readInt(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	return arr
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := readInt(n)
	for i := len(arr); i > 0; i-- {
		fmt.Print(arr[i-1], " ")
	}
}
