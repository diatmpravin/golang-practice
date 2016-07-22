package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func friendList(n, x int, cost []int) string {
	if n == 0 || x == 0 {
		return "NO"
	}

	if cost[0] == x {
		return "YES"
	}

	var start int
	sum := cost[start]

	for i := 1; i < n; i++ {
		sum = sum + cost[i]

		if sum > x {
			for sum > x {
				sum = sum - cost[start]
				start++
			}

		}

		if sum == x {
			return "YES"
			break
		}

	}

	return "NO"
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	for i := 0; i < t; i++ {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())

		s.Scan()
		x, _ := strconv.Atoi(s.Text())

		cost := []int{}
		for i := 0; i < n; i++ {
			s.Scan()
			val, _ := strconv.Atoi(s.Text())
			cost = append(cost, val)
		}

		fmt.Println(friendList(n, x, cost))
	}
}
