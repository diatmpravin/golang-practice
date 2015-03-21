package main

import (
	"fmt"
)

type my struct {
	arr []int
}

func (m *my) dosomething() {
	m.arr = append(m.arr, 1)
	m.arr = append(m.arr, 2)
	m.arr = append(m.arr, 3)
}

func (m my) dosomethingelse() {
	fmt.Println(m.arr)
}

func main() {
	m := new(my)
	m.dosomething()
	m.dosomethingelse()
}
