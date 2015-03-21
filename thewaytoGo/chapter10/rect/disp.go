package main

import "fmt"
import "rect"

type cir struct {
  a bool
  c rect.rectangle
}

func main() {
        rect.Hello() 
	fmt.Println("In Main")

        c := new(cir)

	fmt.Println("Rect====>", c)
}
