// Exercise-14.2
package main

import (
	"fmt"
	"time"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int)
	out <- 2
	go f1(out)
	time.Sleep(5 * 1e9)
}
