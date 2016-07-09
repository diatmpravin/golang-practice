package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s *bufio.Scanner

func main() {
	s = bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	s.Scan()
	k, _ := strconv.Atoi(s.Text())

	arr := make([]int, n)
	E := 100

	for i := 0; i < n; i++ {
		s.Scan()
		tmp, _ := strconv.Atoi(s.Text())
		arr[i] = tmp
	}

	for j := 0; j < n; j += k {
		if arr[j] == 1 {
			E = E - 3
		} else {
			E = E - 1
		}
	}
	fmt.Println(E)	
}