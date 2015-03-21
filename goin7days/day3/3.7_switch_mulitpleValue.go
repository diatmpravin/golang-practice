package main

import "fmt"

func main() {
	var input int = 3

	switch input {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2, 3:
		fmt.Println("Two and Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}
}
