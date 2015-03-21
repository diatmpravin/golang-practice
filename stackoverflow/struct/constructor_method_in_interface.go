package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type ShapeStruct struct {
	Shape
}

type Rectangle struct {
}

func (this *Rectangle) Area() float64 {
	return 42
}

func (s *ShapeStruct) New() Shape {
	return &Rectangle{}
}

func main() {
	a := &ShapeStruct{&Rectangle{}}
	fmt.Println(a.Area())
}
