package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	noOfSpace := n - 1
	noOfStair := 1

	for j := 0; j < n; j++ {
		for i := 0; i < noOfSpace; i++ {
			fmt.Printf(" ")
		}

		for i := 0; i < noOfStair; i++ {
			fmt.Printf("#")
		}

		fmt.Println()
		noOfSpace--
		noOfStair++
	}
}
