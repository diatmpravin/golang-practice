// return.go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello 0")
	firstHello()
	secondHello()
}

func firstHello() {
	var orig = "ABC"
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error %s \n", orig, err)
	}
	fmt.Printf("The interger is %d\n", an)
}

func secondHello() {
	fmt.Println("Hello 2")
}
