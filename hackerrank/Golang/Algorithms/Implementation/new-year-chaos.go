package newyearchaos

import (
	"fmt"
	"sort"
)

func NewYearChaos(arr []int) (count int) {
	fmt.Println(arr)
	counter := len(arr)

	if counter == 0 || counter == 1 {
		return -1
	}

	next := 0
	flag := false

	for i := 0; i < counter-1; i++ {

		if arr[i] != i+1 {
			tmp := arr[i+1]
			arr[i+1] = arr[i]
			arr[i] = tmp

			count++
			flag = true
		}

		if flag {
			next++
			flag = false
		} else {
			next = 0
		}

		if next == 3 {
			return -1
		}

		if next == counter-1 && !sort.IntsAreSorted(arr) {
			return -1
		}
	}

	return
}