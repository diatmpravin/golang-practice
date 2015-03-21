package main

import "fmt"

func main() {
	for i := 0; ; i++ {
		fmt.Printf("The variable i is now: %d\n", i)
		if i > 8 {
			break
		}
	}
}
