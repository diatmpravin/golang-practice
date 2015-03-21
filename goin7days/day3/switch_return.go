package main

import "fmt"

func main() {
	var input int = 0

	switch input {
	case 0:
		fmt.Println("Zero")
		return
	case 1:
		fmt.Println("One")
		return
	default:
		fmt.Println("Unknown Number")
	}
}
