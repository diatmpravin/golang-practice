package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var s *bufio.Scanner

func main() {
	var n int
	fmt.Scan(&n)

	s = bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	indexI := 0
	indexJ := n - 1
	var sum int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			s.Scan()
			v, _ := strconv.Atoi(s.Text())
			if i == j {
				sum += v
			}

			if i == indexI && j == indexJ {
				sum -= v
			}
		}
		indexI++
		indexJ--
	}

	fmt.Println(math.Abs(float64(sum)))
}
