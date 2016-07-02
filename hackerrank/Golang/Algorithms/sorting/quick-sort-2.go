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

func qSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	p := slice[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 1; i < len(slice); i++ {
		if slice[i] < p {
			left = append(left, slice[i])
		} else {
			right = append(right, slice[i])
		}
	}
	left = qSort(left)
	right = qSort(right)
	result := make([]int, 0, len(left) + len(right) + 1)
	for _, n := range(left) {
		result = append(result, n)
	}
	result = append(result, p)
	for _, n := range(right) {
		result = append(result, n)
	}
	printNumbers(result)
	return result
}

func main() {
	s = bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanLines)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	arr := readArr(n)
	qSort(arr)
}