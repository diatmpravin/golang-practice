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
	t, _ := strconv.Atoi(s.Text())

	for i := 0; i < t; i++ {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())

		checks := make([]int, 3)

		s.Scan()
		checks[0], _ = strconv.Atoi(s.Text())

		s.Scan()
		checks[1], _ = strconv.Atoi(s.Text())

		s.Scan()
		checks[2], _ = strconv.Atoi(s.Text())

		sliceInt := make([]int, n)

		for i := 0; i < n; i++ {
			s.Scan()
			tmp, _ := strconv.Atoi(s.Text())
			sliceInt[i] = tmp
		}

		if checkEqual(checks, sliceInt) {
			fmt.Println("She can")

		} else {
			fmt.Println("She can't")
		}
	}
}

func checkEqual(c, s []int) bool {
	flag := true

	for _, checkV := range c {
		for _, v := range s {
			if v%checkV != 0 {
				flag = false
				return flag
			}
		}
	}

	return flag
}
