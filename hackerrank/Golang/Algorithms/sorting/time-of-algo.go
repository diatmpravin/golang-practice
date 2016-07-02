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

	sum := 0

	for x := 1; x < len(arr); x++ {
		tmp := arr[x]
		var i int
		for i = x - 1; i >= 0; i-- {
			if tmp < arr[i] {
				sum += 1
				arr[i + 1] = arr[i]
			} else {
				break
			}
		}
		arr[i+1] = tmp
	}	
	fmt.Println(sum)
}