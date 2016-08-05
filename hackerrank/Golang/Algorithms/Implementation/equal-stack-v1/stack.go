package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return s
}

func maximum(s1, s2, s3 int) (m int) {

	if s1 > s2 && s1 > s3 {
		m = s1
	} else if s2 > s1 && s2 > s3 {
		m = s2
	} else {
		m = s3
	}

	return m
}

func equalStack(arr1, arr2, arr3 []int) int {
	fmt.Println("+++++++++++++ equalStack +++++++++++++++++++++")
	if len(arr1) == 0 && len(arr2) == 0 && len(arr3) == 0 {
		return 0
	}

	if sum(arr1) == sum(arr2) && sum(arr2) == sum(arr3) {
		return sum(arr3)
	}

	max := maximum(sum(arr1), sum(arr2), sum(arr3))

	if max == sum(arr1) {
		newArr := arr1[1:len(arr1)]
		return equalStack(newArr, arr2, arr3)
	} else if max == sum(arr2) {
		newArr := arr2[1:len(arr2)]
		return equalStack(arr1, newArr, arr3)
	} else if max == sum(arr3) {
		newArr := arr3[1:len(arr3)]
		return equalStack(arr1, arr2, newArr)
	}

	return 0
}

func convertToSlice(s []string, n int) []int {
	fmt.Println("++++++++++++++++++++++++>", s)
	arr := []int{}
	var i int
	for i = 0; i < n; i++ {
		// r, err := strconv.Atoi(s[i])
		r, err := strconv.ParseInt(s[i], 10, 64)
		if err != nil {
			fmt.Println("Lenght=====>", len(arr))
			fmt.Println(err)
		}
		arr = append(arr, int(r))
	}

	return arr
}

func main() {
	var n1, n2, n3, i, tmp int

	fmt.Scan(&n1, &n2, &n3)

	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanLines)

	s.Scan()
	arr1 := convertToSlice(strings.Split(s.Text(), " "), n1)
	fmt.Println("########################>", arr1)

	arr2 := make([]int, n2, n2)
	for i = 0; i < n2; i++ {
		fmt.Println(i)
		fmt.Scan(&tmp)
		arr2[i] = tmp
	}
	fmt.Println("+++++++++++++ arr1 +++++++++++++++++++++", arr2)

	arr3 := make([]int, n3)
	for i = 0; i < n3; i++ {
		fmt.Scan(&tmp)
		arr3[i] = tmp
	}

	fmt.Println(equalStack(arr1, arr2, arr3))
}
