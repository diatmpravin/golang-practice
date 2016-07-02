package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var s *bufio.Scanner

func convertToInt(s []string, n int) []int {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		val, _ := strconv.Atoi(s[i])
		arr[i] = val
	}
	return arr
}

func readArr(n int) (arr []int) {
	s.Scan()
	arr = convertToInt(strings.Split(s.Text(), " "), n)
	return
}

func main() {
	s = bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanLines)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	arr := readArr(n)

	p := arr[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] < p {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	for _, n := range(left) {
		fmt.Print(n, " ")
	}
	fmt.Print(p, " ")
	for _, n := range(right) {
		fmt.Print(n, " ")
	}
	fmt.Print("\n")	
}