package main

import "fmt"

func main() {
	var val int = 10

	if val < 10 {
		fmt.Println("Value is less than ten!")
	} else if val == 10 {
		fmt.Println("Value is equal to ten!")
	} else {
		fmt.Println("Vale is greater than tem!")
	}
}
