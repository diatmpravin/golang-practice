package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)
	time.Sleep(10 * 1e9)
}

func sendData(ch chan string) {
	ch <- "pravin"
	ch <- "ankit"
	ch <- "sanyam"
	ch <- "avinaw"
}

func getData(ch chan string) {
	var input string
	for {
		input = <- ch
		fmt.Println(input)
	}
}
