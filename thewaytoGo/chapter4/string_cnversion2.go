package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig = "ABC"
	var an int
	var err error
	an, err = strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error %s \n", orig, err)
		return
	}
	fmt.Printf("The interger is %d\n", an)
}
