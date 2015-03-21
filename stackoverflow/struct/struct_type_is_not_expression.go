package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {
	fmt.Println(salutation.name)
	fmt.Println(salutation.greeting)
}

func main() {
	var s = Salutation{}
	s.name = "Alex"
	s.greeting = "Hi"
	Greet(s)
}
