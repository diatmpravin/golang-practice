package main

import "fmt"

func main() {
	array := [7]int{0, 3, 2, 5, 4, 8, 7}
	for value := range array {
		if value == 5 {
			fmt.Println("Value of boolean expression:", value == 5)
		}
	}

	a, b := 4, 5
	if a < b {
		fmt.Println("a is less than b")
	} else if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("a is equal to b")
	}

	firstName, nickName := "Joey", "Katzen"

	if firstName != nickName {
		fmt.Println("Value of boolean expression:", firstName != nickName)
	}
}
