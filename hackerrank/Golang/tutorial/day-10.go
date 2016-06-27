package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	val := int64(n)
	binary := strconv.FormatInt(val, 2)

	var count, final int
	for _, val := range binary{
		if string(val) == "1" {
			count += 1
		} else {
			if final < count{
				final = count
				count = 0
			}
			count = 0
		}
	}
	if final < count {
		fmt.Println(count)
	} else {
		fmt.Println(final)
	}
}
