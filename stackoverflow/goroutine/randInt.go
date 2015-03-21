package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		amt := time.Duration(rand.Intn(250))
		fmt.Println(amt)
	}
}
