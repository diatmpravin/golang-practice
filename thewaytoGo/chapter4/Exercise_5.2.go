// Exercise_5.2.go
package main

import (
	"fmt"
)

func main() {
	var ans int

	fmt.Printf("Enter day: ")
	fmt.Scanln(&ans)

	switch ans {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thrusday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("You entered wrong no.")
	}
}
