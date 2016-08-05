package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	arr := make([][]rune, n)
	s.Split(bufio.ScanRunes)

	// var row string

	for i := 0; i < n; i++ {
		arr[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			s.Scan()
			fmt.Println(s.Text())
			v, _ := strconv.Atoi(s.Text())
			// fmt.Scan(&row)

			arr[i][j] = rune(v)
		}
	}
	fmt.Println(arr)

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			gap1 := arr[i][j] - arr[i][j-1]
			gap2 := arr[i][j] - arr[i-1][j]
			gap3 := arr[i][j] - arr[i+1][j]
			gap4 := arr[i][j] - arr[i][j+1]

			if gap1 > 0 && gap2 > 0 && gap3 > 0 && gap4 > 0 {
				arr[i][j] = rune('X')
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf(string(arr[i][j]))
		}
		fmt.Println()
	}
}
