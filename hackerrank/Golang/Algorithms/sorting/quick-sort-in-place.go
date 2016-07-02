package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	a := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		a[i], _ = strconv.Atoi(s.Text())
	}

	quicksort(a, 0, len(a)-1)
}

func quicksort(arr []int, lo, lastEl int) {
	if lo < lastEl {
		p := partition(arr, lo, lastEl)
		quicksort(arr, lo, p-1)
		quicksort(arr, p+1, lastEl)
	}
}

func partition(a []int, lo, lastEl int) int {
	pivotIndex := choosePivot(a, lo, lastEl)
	pivotValue := a[pivotIndex]
	a[pivotIndex], a[lastEl] = a[lastEl], a[pivotIndex]
	storeIndex := lo

	for i := lo; i <= lastEl-1; i++ {
		if a[i] < pivotValue {
			a[i], a[storeIndex] = a[storeIndex], a[i]
			storeIndex++
		}
	}

	a[storeIndex], a[lastEl] = a[lastEl], a[storeIndex]

	print(a)

	return storeIndex
}

func choosePivot(arr []int, lo, lastEl int) int {
	return lastEl
}

func sprint(a []int) string {
	str := ""
	for i, v := range a {
		str += fmt.Sprintf("%v", v)
		if i != len(a)-1 {
			str += " "
		}
	}
	return str
}

func print(a []int) {
	str := sprint(a)
	fmt.Printf("%v\n", str)
}