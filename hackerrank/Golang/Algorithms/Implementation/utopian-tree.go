package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculateHight(n int) int {
	l := 1

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			l++
		} else {
			l *= 2
		}
	}

	return l
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	for i := 0; i < t; i++ {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())

		fmt.Println(calculateHight(n))
	}
}
