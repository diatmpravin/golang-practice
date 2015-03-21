package main

import "fmt"

func main() {
	var first int = 10

	if first <= 0 {
		fmt.Println("The first is less and equal to 0")
	} else if first > 0 && first < 5 {
		fmt.Println("The first is greater than 0 and less than 5")
	} else {
		fmt.Println("The first is greate than 5")
	}

	if cond := 10; cond < 10 {
		fmt.Println("Cond is less than 10")
	} else {
		fmt.Println("Cond is greater than 10")
	}
}
