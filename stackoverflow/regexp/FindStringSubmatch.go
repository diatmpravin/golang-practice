package main

import (
	"fmt"
	"regexp"
)

var myExp = regexp.MustCompile(`(?P<first>\d+)\.(\d+).(?P<second>\d+)`)

func main() {
	fmt.Printf("%+v", myExp.FindStringSubmatch("1234.5678.9"))
}
