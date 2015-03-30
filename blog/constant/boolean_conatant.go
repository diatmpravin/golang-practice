package main

import (
	"fmt"
)

func main() {
	type MyBool bool
    const True = true
    const TypedTrue bool = true
    var mb MyBool
    mb = true      // OK
    mb = True      // OK
    // mb = TypedTrue // Bad
    fmt.Println(mb)
}