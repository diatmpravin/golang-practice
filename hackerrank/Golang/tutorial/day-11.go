package main

import (
	"fmt"
)

var highscore = -1000

func check(m, n int, a [6][6]int) int {
	var res int = 0
	var tmp int = n
	if m < 4 && n < 4 {
		for x := 0; x < 3; x++ {
			if x == 1 {
				res += a[tmp][m+1]
			} else {
				res += a[tmp][m] + a[tmp][m+1] + a[tmp][m+2]
			}
			tmp += 1
		}
		return res
	}
	return -100
}

func main() {
	var arr = [6][6]int{}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[j]); j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	var s int
	for x := 0; x < len(arr); x++ {
		for y := 0; y < len(arr[x]); y++ {
			s = check(x, y, arr)
			if s > highscore {
				highscore = s
			}
		}
	}
	fmt.Println(highscore)
}