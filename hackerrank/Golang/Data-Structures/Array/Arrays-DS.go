package main

import (
	"fmt"
)

var n int

func readInt(n int)([]int) {
	arr := make([]int, n)
	for i := range arr{
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			fmt.Println(err)
		}
	}

	return arr
}

func main() {
	
	fmt.Scanln(&n)
	arr := readInt(n)

	for i := range arr{
		fmt.Print(arr[n-(i+1)], " ")
	}
}