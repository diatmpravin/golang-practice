package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s *bufio.Scanner

func checkEven(arr []int) (flag bool) {
	for i := 0; i < len(arr); i++ {
		if arr[i] % 2 != 0 {
			flag = true
			break 
		}
	}

	if flag {
		return false
	} else {
		return true
	}
}

func fairRations(arr []int, n, c int ) int {
	flag, done := false, false
	for i := 0; i < n-1; i++ {
		if arr[i]%2 != 0 {
			arr[i], arr[i+1] = arr[i] + 1, arr[i+1] + 1
			c +=2
			flag = true
			if checkEven(arr) {
				done = true
				break
			} 
		}
	}

	if flag && done {
		return c	
	} else {
		return -1
	}

	
}

func main() {
	s = bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		tmp, _ := strconv.Atoi(s.Text())

		arr[i] = tmp
	}

	var c int
	c = fairRations(arr, n, c)

	if c >= 0 {
		fmt.Println(c)
	} else {
		fmt.Println("NO")
	}
}