package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var x *bufio.Scanner

func convertToInt(s []string) []int {
	arr := make([]int, 3)
	for i := 0; i < 3; i++ {
		v, _ := strconv.Atoi(s[i])
		arr[i] = v
	}
	return arr
}

func main() {
	x = bufio.NewScanner(bufio.NewReader(os.Stdin))
	x.Split(bufio.ScanWords)

	x.Scan()
	t, _ := strconv.Atoi(x.Text())

	for i := 0; i < t; i++ {
		x.Split(bufio.ScanLines)
		x.Scan()
		nms := convertToInt(strings.Split(x.Text(), " "))

		n := nms[0]
		m := nms[1]
		s := nms[2]

		p := m%n

		// if m < n {
		// 	fmt.Println(m+s-1)
		// }

		// if m > n {
		// 	fmt.Println(p+s-1)
		// }
		var id int
		if p == 0 {
			p = n
		}

		if s+p-1 > n {
			id = (s-1) - (n-p)
		} else {
			id = (s-1) + p
		}

		fmt.Println(id)
	}
}
