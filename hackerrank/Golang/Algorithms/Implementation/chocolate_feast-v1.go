package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convToInt(s []string) []int {
	i := []int{}

	for _, v := range s {
		val, _ := strconv.Atoi(v)
		i = append(i, val)
	}
	return i
}

func calculateChocolate(n, c int) (ch, r int) {
	ch = n / c
	r = n % c
	return
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	var ch int

	for i := 0; i < t; i++ {
		s.Split(bufio.ScanLines)

		s.Scan()
		ncm := convToInt(strings.Split(s.Text(), " "))

		n := ncm[0]
		c := ncm[1]
		m := ncm[2]

		ch, _ = calculateChocolate(n, c)

		for i := ch; i >= m; {
			r2ch, r := calculateChocolate(i, m)
			ch = ch + r2ch
			i = r + r2ch	
		}

		fmt.Println(ch)
	}
}
