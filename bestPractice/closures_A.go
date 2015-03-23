package main

import (
	"fmt"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	for ix := range values {
		func() {
			fmt.Println(ix)
		}()
	}
}
