package main

import "fmt"

func main() {
	var t, x, y, z, b, w int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&b, &w, &x, &y, &z)
		var result int

		if x > y+z {
			x = y + z
		} else if y > x+z {
			y = x + z
		}

		result = x*b + y*w
		fmt.Println(result)
	}
}
