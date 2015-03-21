// 6.1-greeting.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("I am before Greeting")
	greeting()
	fmt.Println("I am after Greeting")
}

func greeting() {
	fmt.Println("I am in greeting")
}
