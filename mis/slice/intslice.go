package main

import (
	"fmt"
	"sort"
)

func main() {
	keys := []int{3, 2, 8, 1}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	fmt.Println(keys)
}
