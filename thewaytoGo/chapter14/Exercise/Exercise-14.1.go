// Exercise 14.1.go
package main

import (
	"fmt"
	"time"
)

func f(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	ch := make(chan int)
	go f(ch)
	ch <- 3
	time.Sleep(time.Second)
}
