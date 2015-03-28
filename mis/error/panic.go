package main

import (
	"fmt"
)

func letUsPanic() {
	defer func() {
		fmt.Println("This will be called")
	}()
	panic("boom!")
	fmt.Println("This will never be called")
}

// func letUsPanic() {
// 	defer func() {
// 		if e := recover(); e != nil {
// 			// e is the interface{} typed-value we passed to panic()
// 			fmt.Println("Whoops: ", e) // Prints "Whoops: boom!"
// 		}
// 	}()
// 	panic("boom!")
// 	fmt.Println("This will never be called")
// }

func main() {
	letUsPanic()
}
