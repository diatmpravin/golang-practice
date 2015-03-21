package main

import "fmt"

func main() {
	for i := 0; i < 2; i++ {
		for j := 5; j < 8; j++ {
			if j > 6 {
				break
			}
		fmt.Println("Value of I and J is:", i, j)
		}
	}
}
