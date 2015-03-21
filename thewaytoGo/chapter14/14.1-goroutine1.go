package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main")
	go longWait()
	go shortWait()
	fmt.Println("About to sleep")
	time.Sleep(10 * 1e9)
	fmt.Println("end main")
}

func longWait() {
	fmt.Println("in longwait")
	time.Sleep(5 * 1e9)
	fmt.Println("end longwait")
}

func shortWait() {
	fmt.Println("In shortWait")
	time.Sleep(2 * 1e9)
	fmt.Println("End short Wait")
}
