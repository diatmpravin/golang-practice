package main

import "fmt"

func main() {
	for i, j := 0, 3; i <= j; i, j = i+1, j-1 {
		fmt.Println("This is the I and J value respectively:", i, j)
	}
}
