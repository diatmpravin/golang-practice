package main

import (
	"container/list"
	"fmt"
)

func main() {
	alist := list.New()

	fmt.Println(alist.Len()) // list size before

	alist.PushBack("a")
	alist.PushBack("b")
	alist.PushBack("c")

	fmt.Println(alist.Len()) // list size after

}
