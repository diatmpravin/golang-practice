package main

import (
	"fmt"
	"time"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	for ix := range values {
		go func() {
			fmt.Println(ix)
		}()
	}
	time.Sleep(5 * time.Second)
}
