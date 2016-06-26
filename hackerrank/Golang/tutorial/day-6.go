package main

import (
	"bytes"
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	if T >= 0 && T <= 10 {
		for i := 0; i < T; i++ {
			var S string
			fmt.Scan(&S)
			length := len(S)
			var even, odd bytes.Buffer

			if length >= 2 && length <= 10000 {
				for key, value := range S {
					if key == 0 || key%2 == 0 {
						even.WriteString(string(value))
					} else {
						odd.WriteString(string(value))
					}
				}
			}
			fmt.Println(even.String(), odd.String())
		}
	}
}
