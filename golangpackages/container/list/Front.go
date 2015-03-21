package main

import (
	"container/list"
	"fmt"
)

func main() {
	alist := list.New()

	fmt.Println(alist)

	alist.PushBack("a")
	fmt.Println(alist)

	alist.PushBack("b")
	fmt.Println(alist)

	alist.PushBack("c")
	fmt.Println(alist)

	e := alist.Front()
	fmt.Println(e.Value)
}
