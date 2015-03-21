// Question 7.1.go
package main

import (
	"fmt"
)

func main() {
	a := [...]string{"a", "b", "c", "d", "e"}
	for i, v := range a {
		fmt.Printf("Value at index %d is %s\n", i, v)
	}

}
