package main

func findElement(s []int, index int) int {
	return s[index]
}

func circularArrayRotation(a []int, k int) []int {
	for i := 0; i < k; i++ {
		newSlice := make([]int, len(a))
		for index, _ := range a {
			if index == len(a)-1 {
				newSlice[0] = a[index]
			} else {
				newSlice[index+1] = a[index]
			}
		}
		a = newSlice
	}

	return a
}
