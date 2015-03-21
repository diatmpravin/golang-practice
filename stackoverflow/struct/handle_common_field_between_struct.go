package main

import "fmt"

type A struct {
	X int
	Y int
}

func (a *A) Sum() int {
	return a.X + a.Y
}

type B struct {
	*A
	Z int
}

func main() {
	a := &A{1, 2}
	b := &B{&A{3, 4}, 5}
	fmt.Println(a.Sum(), b.Sum())
}
