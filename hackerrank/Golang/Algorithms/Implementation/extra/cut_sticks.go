package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	exits := n

	sticks := make([]int, n)

	cut := 32768
	for i := 0; i < n; i++ {
		fmt.Scan(&sticks[i])
		if sticks[i] < cut {
			cut = sticks[i]
		}
	}
	fmt.Println(exits)
J:
	for {
		short := 32768
		for i := 0; i < n; i++ {
			if sticks[i] > 0 {
				sticks[i] -= cut
				if sticks[i] == 0 {
					exits--
					if exits == 0 {
						break J
					}
				} else {
					if sticks[i] < short {
						short = sticks[i]
					}
				}
			}
		}

		fmt.Println(exits)
		cut = short
	}
}