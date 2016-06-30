package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanLines)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	for i := 0; i < n; i++ {
		s.Scan()
		nk := convToInt(strings.Split(s.Text(), " "))

		s.Scan()
		attendance := convToInt(strings.Split(s.Text(), " "))
		onTime := 0

		for _, v := range attendance {
			if v <= 0 {
				onTime++
			}
		}

		if onTime < nk[1] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")

		}
	}
}

func convToInt(s []string) []int {
	i := []int{}

	for _, v := range s {
		val, _ := strconv.Atoi(v)
		i = append(i, val)
	}
	return i
}
