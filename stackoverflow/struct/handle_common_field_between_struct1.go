package main

import "fmt"

type A struct {
	one int
	two int
}

type B struct {
	A
	three int
}

func(x A) sum () int {
	return x.one + x.two
}

func main() {
	a := A{1,2}
	b := B{A{3,4}, 5}
	
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
	
	fmt.Println("Summation of A:", a.sum())
	fmt.Println("Summation of A:", b.sum())
}
