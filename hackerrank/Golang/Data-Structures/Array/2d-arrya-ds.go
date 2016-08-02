package main

import "fmt"

func main() {
	a := [][]int{}

	for i := 0; i < 6; i++ {
		buf := []int{}

		for j := 0; j < 6; j++ {
			var t int

			fmt.Scan(&t)

			buf = append(buf, t)
		}

		a = append(a, buf)
	}

	maxSum := -100
	for i := 0; i < 4; i++ {
		for j := 1; j < 5; j++ {
			curSum := a[i][j-1] + a[i][j] + a[i][j+1] +
				a[i+1][j] +
				a[i+2][j-1] + a[i+2][j] + a[i+2][j+1]

			if curSum > maxSum {
				maxSum = curSum
			}
		}
	}

	fmt.Println(maxSum)
}
