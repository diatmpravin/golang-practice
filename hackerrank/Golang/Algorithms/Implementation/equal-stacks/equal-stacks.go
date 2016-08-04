package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func primeSlice(a, b, c int64) (p int64) {
	if a > b && a > c {
		p = a
	} else if b > a && b > c {
		p = b
	} else {
		p = c
	}

	return
}

func calculateSum(s []int64) (sum int64) {
	for _, v := range s {
		sum += v
	}

	return
}

func equalStack(slice1, slice2, slice3 []int64) int64 {
	s1 := calculateSum(slice1)
	s2 := calculateSum(slice2)
	s3 := calculateSum(slice3)

	var c1, c2, c3 int64

	for {
		if s1 == s2 && s2 == s3 {
			break
		}

		max := primeSlice(s1, s2, s3)
		// fmt.Println(max)
		// fmt.Println(s1, s2, s3)

		if max == 0 {
			return max
		}

		if max == s1 {
			s1 -= slice1[c1]
			c1++
		} else if max == s2 {
			s2 -= slice2[c2]
			c2++
		} else if max == s3 {
			s3 -= slice3[c3]
			c3++
		}

	}

	return s1
}

func convertToSlice(s []string, n int64) []int64 {
	arr := make([]int64, n)
	var i int64
	for i = 0; i < n; i++ {
		r, err := strconv.Atoi(s[i])
		if err != nil {
			panic(err)
		}
		arr[i] = int64(r)
	}

	return arr
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanLines)

	s.Scan()
	n1n2n3 := convertToSlice(strings.Split(s.Text(), " "), 3)

	s.Scan()
	slice1 := convertToSlice(strings.Split(s.Text(), " "), n1n2n3[0])

	s.Scan()
	slice2 := convertToSlice(strings.Split(s.Text(), " "), n1n2n3[1])

	s.Scan()
	slice3 := convertToSlice(strings.Split(s.Text(), " "), n1n2n3[2])

	fmt.Println(equalStack(slice1, slice2, slice3))
}
