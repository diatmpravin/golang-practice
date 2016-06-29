package main

import "fmt"

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
	var swaps, tmp, numberOfSwaps int

	for i := 0; i < n; i++ {
		swaps = 0
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				swaps += 1
				numberOfSwaps += 1
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
		if swaps == 0 {
			break
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\n", numberOfSwaps)
	fmt.Printf("First Element: %d\n", arr[0])
	fmt.Printf("Last Element: %d\n", arr[n-1])
}
