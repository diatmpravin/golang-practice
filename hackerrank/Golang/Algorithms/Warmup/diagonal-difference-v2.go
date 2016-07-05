package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"math"
)

var s *bufio.Scanner

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([][]int, n)

	s = bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			s.Scan()
			v, _ := strconv.Atoi(s.Text())
			arr[i][j] = v			 			
		}
	}

	indexI := 0
	indexJ := n-1
	var right, left int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				right += arr[i][j]
			}

			if i == indexI && j == indexJ {
				left += arr[i][j]
			}		
		}
		indexI++
		indexJ--
	}

	fmt.Println(math.Abs(float64(right-left)))
}