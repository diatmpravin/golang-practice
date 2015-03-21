package main

  import (
    "container/list"
    "fmt"
  )

  func main() {
    alist := list.New()
    alist.PushBack("a")
    alist.PushBack("b")
    alist.PushBack("c")

    e := alist.Back()
    fmt.Println(e.Value)
 }