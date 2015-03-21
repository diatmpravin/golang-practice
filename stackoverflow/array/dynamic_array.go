package main

import "fmt"

func main() {
	var elems int
	sum := 0

	fmt.Print("Number of elements? ")

	fmt.Scan(&elems)

	var array = make([]int, elems)

	for i := 0; i < elems; i++ {
		fmt.Printf("%d . Number? ", i+1)
		fmt.Scan(&array[i])
		sum += array[i]
	}

	fmt.Printf("Sum: %d\n", sum)	
}
