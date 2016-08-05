package main

import "fmt"

func main() {
	var cases int
	fmt.Scan(&cases)

	for i := 0; i < cases; i++ {
		var stones, a, b int
		fmt.Scan(&stones, &a, &b)

		if stones == 0 {
			fmt.Println(0)
		} else if a == b {
			fmt.Printf("%d\n", a*(stones-1))
		} else {

			if a > b {
				a, b = b, a
			}

			for j := 0; j < stones; j++ {
				score := b*j + a*(stones-1-j)
				fmt.Printf("%d ", score)
			}
			fmt.Println()
		}
	}
}
