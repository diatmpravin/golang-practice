package main

import (
	"fmt"
	//"strings"
)

type Test struct {
	someStrings []string
}

func (this *Test) AddString(s string) {
	this.someStrings = append(this.someStrings, s)
	this.Count() // will print "1"
}

func (this Test) Count() {
	fmt.Println(len(this.someStrings))
}

func main() {
	var test Test
	test.AddString("testing")
	test.Count() // will print "0"
}
