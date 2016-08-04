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

	s.Scan()
	k, _ := strconv.Atoi(s.Text())

	var index, count int

	for i := 1; i <= n; i++ {
		index++

		s.Scan()
		q, _ := strconv.Atoi(s.Text())		
		
		for j := 0; j < q; j++ {
			if j % k == 0 && j != 0 {
				index++
			}

			if index == j+1 {
				count++
			}
		}
	}
	fmt.Println(count)

}