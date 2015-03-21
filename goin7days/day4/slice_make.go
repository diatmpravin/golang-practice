package main

import "fmt"

func main() {
	slice1 := make([]int, 5)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}

	fmt.Println(slice1)
}
