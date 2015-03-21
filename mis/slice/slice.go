package main

import (
	"fmt"
	"sort"
)

func main() {
	example := []int{1, 25, 3, 5, 4}
	sort.Ints(example)
	fmt.Println(example)
}
