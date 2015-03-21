package main

import "fmt"

func main() {
	var input int = 0

	switch input {
	case 0:
		fmt.Println("Zero")
		fallthrough
	case 1:
		//{
		fmt.Println("One")
		fmt.Println("I am One.! ")
		//}
		//fallthrough
	default:
		fmt.Println("Unknown Number")
	}
}
