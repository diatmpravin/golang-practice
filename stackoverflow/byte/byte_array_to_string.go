package main

import "fmt"

func main() {
	var bytes [5]byte
	bytes[0] = 65

	fmt.Println(string(bytes[:]))
	fmt.Println(bytes[:])
}
