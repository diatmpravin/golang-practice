package main

import (
	"fmt"
)

// Constant
const thing = 5       // untype
const things2 int = 5 // typed -  only used in int expression

// Enum
const (
	One   = 1
	Two   = 2
	Three = 4
)

// Iota
const (
	Five   = 5 << iota // 5 (i.e. 5 << 0)
	Six              // 10 (i.e. 5 << 1)
	Seven             // 20 (i.e. 5 << 2)
)

func main() {
	fmt.Println("thing", thing)
	fmt.Println("things2", things2)

	fmt.Println("One", One)
	fmt.Println("Two", Two)

	fmt.Println("Five", Five)
	fmt.Println("Six", Six)
	fmt.Println("Seven", Seven)
}
