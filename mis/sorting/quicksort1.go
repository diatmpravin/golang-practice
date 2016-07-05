package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s *bufio.Scanner

func printS(s []int) {
	for _, v := range s {
		fmt.Print(v, " ")
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

	left := []int{}
	right := []int{}
	equal := intS[0]

	for x := 1; x < n; x++ {
		if intS[x] > equal {
			right = append(right, intS[x])
		} else if intS[x] < equal {
			left = append(left, intS[x])
		}
	}

	printS(left)
	fmt.Print(equal, " ")
	printS(right)
}
