package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	// "strings"
)

// func convertToSlice(s []string, n int) []int {
// 	fmt.Println("Input:", s)
// 	arr := []int{}
// 	for i := 0; i < n; i++ {
// 		r, err := strconv.Atoi(s[i])
// 		if err != nil {
// 			fmt.Println("Lenght=====>", len(arr))
// 			fmt.Println(err)
// 		}
// 		arr = append(arr, r)
// 	}

// 	return arr
// }

func main() {
	n := 58668

	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanRunes)

	// arr1 := convertToSlice(strings.Split(s.Text(), " "), n)
	for i := 0; i < n; i++ {
		s.Scan()
		if i == 1404 {
			fmt.Println(i, s.Text())
		}

		if i == 1405 {
			fmt.Println(i, s.Text())
		}

		if i == 1406 {
			fmt.Println(i, s.Text())
		}

		if i == 1407 {
			fmt.Println(i, s.Text())
		}

		if i == 1408 {
			fmt.Println(i, s.Text())
		}

		if i == 1409 {
			fmt.Println(i, s.Text())
		}

		if i == 1430 {
			fmt.Println("+++++ breaking ++++")
			break
		}
	}

	// fmt.Println(arr1)
}
