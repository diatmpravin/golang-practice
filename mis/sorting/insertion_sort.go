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

	tmp := intS[n-1]

	var i int
	for i = len(intS) - 2; i >= 0; i-- {
		if tmp < intS[i] {
			intS[i + 1] = intS[i]
		} else {
			break
		}
		printSlice(intS)
	}
	intS[i + 1] = tmp
	printSlice(intS)
}
