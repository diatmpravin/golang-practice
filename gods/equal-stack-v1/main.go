package main

import (
	"fmt"
)

func main() {
	n := 58668
	var tmp int

	arr1 := []int{}
	for i := 0; i < n; i++ {
		fmt.Println("=========================", i)
		fmt.Scan(&tmp)
		arr1 = append(arr1, tmp)
	}

	fmt.Println("==========================", arr1)
}
