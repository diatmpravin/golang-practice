package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findProduct(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	p := arr[0]

	if len(arr) >= 2 {
		for i := 1; i < len(arr); i++ {
			p = p * arr[i] % 1000000007
		}
	}

	return p
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}

	sliceInt := []int{}
	for i := 0; i < n; i++ {
		s.Scan()
		v, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}

		sliceInt = append(sliceInt, v)
	}

	fmt.Println(findProduct(sliceInt))
}
