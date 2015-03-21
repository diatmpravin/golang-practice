package main

import "fmt"

func main() {
	var day int
	fmt.Printf("Enter day> ")
	fmt.Scanln(&day)

	switch day {

	case 1:
		fmt.Println("Sun")
	case 2:
		fmt.Println("Mon")
	case 3:
		fmt.Println("Tue"); fallthrough;
	case 4:
		fmt.Println("Wed")
	case 5:
		fmt.Println("Thu")
	case 6:
		fmt.Println("Fri")
	case 7:
		fmt.Println("Sat")
	default:
		fmt.Println("You have entered wrong days")

	}

}
