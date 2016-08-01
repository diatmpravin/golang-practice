package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
	// "strings"
)

func primeSlice(a, b, c int) (p int) {
	if a < b && b < c {
		p = a
	} else if b < a && b < c {
		p = b
	} else {
		p = c
	}

	return
}

func calculateSum(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}

	return
}

func equalStack(slice1, slice2, slice3 []int) (n int) {
	s1 := calculateSum(slice1)
	s2 := calculateSum(slice2)
	s3 := calculateSum(slice3)

	pSlice := primeSlice(s1, s2, s3)
	fmt.Println(pSlice)

	for {
		
	}

	return
}

// func convertToSlice(s []string, n int) []int {
// 	arr := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		r, err := strconv.Atoi(s[i])
// 		if err != nil {
// 			panic(err)
// 		}
// 		arr[i] = r
// 	}

// 	return arr
// }

// // func checkStack(a, b, c int) bool {
// // 	if a == b && b == c {
// // 		return true
// // 	}

// // 	return false
// // }

// func main() {
// 	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
// 	s.Split(bufio.ScanLines)

// 	s.Scan()
// 	n1n2n3 := convertToSlice(strings.Split(s.Text(), " "), 3)

// 	s.Scan()
// 	slice1 := convertToSlice(strings.Split(s.Text(), " "), n1n2n3[0])

// 	s.Scan()
// 	slice2 := convertToSlice(strings.Split(s.Text(), " "), n1n2n3[1])

// 	s.Scan()
// 	slice3 := convertToSlice(strings.Split(s.Text(), " "), n1n2n3[2])

// 	s1 := calculateSum(slice1)
// 	s2 := calculateSum(slice2)
// 	s3 := calculateSum(slice3)

// 	pSlice := primeSlice(s1, s2, s3)

// 	if pSlice == s1 == s2 == s3 {
// 		fmt.Println(s1, s2, s3, pSlice)
// 	}
// }
