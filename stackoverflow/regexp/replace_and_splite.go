package main

import (
	"fmt"
	"regexp"
)

func main() {
	rp := regexp.MustCompile("([a-z]+) ([a-z]+)")
	fmt.Println(rp.ReplaceAllString("abc def ghi", "$2 $1"))

	rp1 := regexp.MustCompile("a")
    i := rp1.Split("zzzzazzzzz", -1) // ["zzzz", "zzzz"]
    fmt.Println(i)
}
