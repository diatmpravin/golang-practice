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

	for i := 0; i < n; i++ {
		arr[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			s.Scan()
			v, _ := strconv.Atoi(s.Text())

			arr[i][j] = rune(v)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && i < n && j > 0 && j < n {
				fmt.Printf("X")
			} else {
				fmt.Printf(string(arr[i][j]))
			}
		}
		fmt.Println()
	}
}
