package main

import "fmt"

func main() {
	var dayReturn, monthReturn, yearReturn, dayExpected, monthExpected, yearExpected int
	fmt.Scan(&dayReturn)
	fmt.Scan(&monthReturn)
	fmt.Scan(&yearReturn)

	fmt.Scan(&dayExpected)
	fmt.Scan(&monthExpected)
	fmt.Scan(&yearExpected)

	if dayReturn, dayExpected >= 1 {
		fmt.Println("Hello")
	} 
}