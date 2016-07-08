package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMimOfSlice(min []int) (r int) {
	r = -1
	if len(min) == 0 {
		return r
	} else {
		r = min[0]
		for _, v := range min {
			if r > v {
				r = v
			}
		}
		return r
	}
	return r
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanLines)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	s.Scan()
	intS := convToInt(strings.Split(s.Text(), " "), n)

	min := []int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if intS[i] == intS[j] {
				min = append(min, j-i)
			}
		}
	}

	fmt.Println(getMimOfSlice(min))
}

func convToInt(s []string, n int) []int {
	a := make([]int, n)

	for i := 0; i < n; i++ {
		val, _ := strconv.Atoi(s[i])
		a[i] = val
	}

	return a
}
