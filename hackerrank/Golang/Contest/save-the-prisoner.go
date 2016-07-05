package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var s *bufio.Scanner

func convertToInt(s []string) []int {
	fmt.Println(s)
	arr := make([]int, 3)
	for i := 0; i < 3; i++ {
		v, _ := strconv.Atoi(s[i])
		arr[i] = v
	}
	return arr
}

func main() {
	s = bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	for i := 0; i < t; i++ {
		s.Split(bufio.ScanLines)
		s.Scan()
		mns := convertToInt(strings.Split(s.Text(), " "))

		fmt.Println(mns)
	}
}
