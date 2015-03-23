package main

import (
	"fmt"
)

const (
	one   = iota // 0
	two          // 1
	three        // 2
	four         // 3
)

// Here iota used as expressions, storing the resulting value in the constant.
const (
	one1   = 1 << iota // 1
	two1               // 2
	four1              // 4
	eight1             // 8
)

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

//The iota increments on the next line, rather than as soon as it gets referenced.
const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry, Fig
)

func main() {
	fmt.Println("one", one)
	fmt.Println("two", two)
	fmt.Println("three", three)
	fmt.Println("four", four)

	fmt.Println("")

	fmt.Println("one", one1)
	fmt.Println("two", two1)
	fmt.Println("four1", four1)
	fmt.Println("eight1", eight1)

	fmt.Println("")

	fmt.Println("KB", KB)
	fmt.Println("MB", MB)
	fmt.Println("GB", GB)

	fmt.Println("")

	fmt.Println("Apple", Apple)
	fmt.Println("Banana", Banana)
	fmt.Println("Cherimoya", Cherimoya)
	fmt.Println("Durian", Durian)
	fmt.Println("Elderberry", Elderberry)
	fmt.Println("Fig", Fig)

}
