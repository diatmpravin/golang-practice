package main

import (
	"fmt"
	"regesp"
)

func main() {
	rp := regexp.MustCompile("([a-z]+) ([a-z]+)")
	rp.ReplaceAllString("abc def ghi", "$2 $1")
}
