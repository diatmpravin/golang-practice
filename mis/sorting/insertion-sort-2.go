package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s *bufio.Scanner

func printSlice(s []int) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Print("\n")
}

func insertionSort(intS []int) {
	for j := 1; j < len(intS); j++ {
		for i := 1; i < len(intS); i++ {
			if intS[i] < intS[i-1] {
				tmp := intS[i]
				intS[i] = intS[i-1]
				intS[i-1] = tmp
			}
		}
		printSlice(intS)
	}
}

func main() {
	s = bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}

	intS := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		intS[i], err = strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}
	}

	insertionSort(intS)
}
