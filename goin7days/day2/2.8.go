package main

import "fmt"

func main() {
	marks := [5]int{85, 91, 89, 98, 92}

	for i, v := range marks {
		fmt.Println("Value at index ", i, "is:", v)
	}
}
