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

func printNumbers(slice []int) {
	for _, n := range(slice) {
		fmt.Print(n, " ")
	}
	fmt.Print("\n")
}

func main() {
	s = bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanLines)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	arr := readArr(n)
	
	tmp := arr[n-1]

	var i int
	for i = len(arr) - 2; i >= 0; i-- {
		if tmp < arr[i] {
			arr[i + 1] = arr[i]
		} else {
			break
		}
		printNumbers(arr)
	}
	arr[i + 1] = tmp
	printNumbers(arr)	
}
