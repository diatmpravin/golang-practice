package main

import "fmt"

func main() {
	var n, i, pos, neg, zeros, temp float64
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&temp)

		if temp > 0 {
			pos += 1
		} else if temp < 0 {
			neg += 1
		} else if temp == 0 {
			zeros += 1
		}
	}

	fmt.Println(pos, neg, zeros)
	fmt.Printf("%.6f \n%.6f\n %.6f \n", pos/n, neg/n, zeros/n)
}
