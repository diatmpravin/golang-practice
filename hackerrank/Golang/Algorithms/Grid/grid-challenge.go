package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	s.Split(bufio.ScanLines)

	for k := 0; k < t; k++ {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())

		g := make([][]int, n)
		for i := 0; i < n; i++ {
			s.Scan()
			line := s.Text()
			d := make([]int, n)
			for j := 0; j < n; j++ {
				d[j] = int(line[j])
			}
			sort.Ints(d)
			g[i] = d
		}

		flag := true
		for i := 0; i < n; i++ {
			for j := 1; j < n; j++ {
				if g[j-1][i] > g[j][i] {
					flag = false
					break
				}
			}
		}

		if flag {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
