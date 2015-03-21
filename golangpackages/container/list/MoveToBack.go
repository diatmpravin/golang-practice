package main

import (
	"container/list"
	"fmt"
)

func main() {
	alist := list.New()
	alist.PushBack("a")
	alist.PushBack("b")
	e := alist.PushBack("c")
	alist.PushBack("d")

	alist.MoveToBack(e)

	for e := alist.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // print out the elements
	}

}
