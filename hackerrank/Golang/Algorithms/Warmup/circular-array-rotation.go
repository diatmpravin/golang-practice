package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
)

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

func main() {
	var n, k, q int

	fmt.Scan(&n)	
	fmt.Scan(&k)	
	fmt.Scan(&q)

	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		v, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}

		slice[i] = v
	}
	newSlice := circularArrayRotation(slice, k)
	
	for i := 0; i < q; i++ {
		s.Scan()
		v, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}

		fmt.Println(findElement(newSlice, v))		
	}

}
