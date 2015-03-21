package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(10)

	fmt.Printf("%#v", r)

	fmt.Printf("Ring size is : %d\n", r.Len())

}
