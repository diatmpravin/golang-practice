package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s *bufio.Scanner

func readArray(n int) []string{
	strArray := make([]string, n)

	for i := 0; i < n; i++ {
		s.Scan()
		strArray[i] = s.Text()
	}

	return strArray
}

func countQuery(q string, d []string) (count int) {
	for _, v := range d {
		if v == q {
			count++
		}
	}

	return
}

func main() {
	s = bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanLines)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	dataArry := readArray(n)

	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	for i := 0; i < t; i++ {
		s.Scan()
		fmt.Println(countQuery(s.Text(), dataArry))
	}
}
