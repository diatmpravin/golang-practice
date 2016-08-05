package main

import "fmt"

func main() {
	var t, n, c, m int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&n, &c, &m)

		amount := n / c
		wrapper := amount

		for {
			if wrapper >= m {
				cho := wrapper / m
				amount += cho
				wrapper = cho + wrapper - cho*m
			} else {
				fmt.Println(amount)
				break
			}
		}
	}
}
