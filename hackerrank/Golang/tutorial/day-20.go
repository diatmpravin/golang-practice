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
	var swaps, tmp int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				swaps += 1
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\n", swaps)
	fmt.Printf("First Element: %d\n", arr[0])
	fmt.Printf("Last Element: %d\n", arr[n-1])
}
