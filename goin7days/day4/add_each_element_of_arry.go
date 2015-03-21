package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	var total int

	for _, v := range arr {
		total = total + v
	}

	fmt.Println("Total:", total)
}
