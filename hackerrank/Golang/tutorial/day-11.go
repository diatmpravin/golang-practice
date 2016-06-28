package main

import "fmt"

func main() {
	A := [6][6]int{}

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			var val int
			fmt.Scan(&val)
			A[i][j] = val
		}
	}
	fmt.Println(A)
	fmt.Println()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			
		}
	}
}
