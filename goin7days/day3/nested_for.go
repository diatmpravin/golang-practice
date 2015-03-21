package main

import "fmt"

func main() {
	firstValue, secondValue := 0, 10

	for i = firstValue; i < 2; i++ {
		for j = secondValue; j < 12; j++ {
			fmt.Println(j)
		}
	}

}
