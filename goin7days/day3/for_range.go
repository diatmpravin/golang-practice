package main

import "fmt"

func main() {
	str := "GopherConIndia!"
	for pos, char := range str {
		fmt.Printf("Character on position %d is %c: \n", pos, char)
	}
}
