package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"bca", "abc", "acb"}
	sort.Strings(s)
	fmt.Println(s)
}
