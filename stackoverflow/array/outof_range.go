package main

import (
	"fmt"
)

type Player struct {
	name string
}

type Board struct {
	Tboard  [9]string
	Player1 Player
	Player2 Player
}

func makeBoard() *Board {
	b := &Board{Tboard: [9]string{}}
	for x := 0; x < len(b.Tboard); x++ {
		b.Tboard[x] = "E"
		fmt.Println(b.Tboard[x])
	}
	fmt.Println(len(b.Tboard)) // => 9
	fmt.Print("Error: ")
	fmt.Println(b)        // => Error: %!v(PANIC=runtime error: index out of range)
	fmt.Println(b.Tboard) // => [[ ] [ ] [ ] [ ] [ ] [ ] [ ] [ ] [ ]]
	return b
}

func main() {
	fmt.Println(makeBoard())
}
