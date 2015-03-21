package main

import "fmt"

func main() {
	var input int = 0

	switch input {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	default:
		fmt.Println("Unknown Number")
	}
}
