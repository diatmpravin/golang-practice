package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[int]string)
	m[1] = "a"
	m[2] = "c"
	m[0] = "b"

	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}
