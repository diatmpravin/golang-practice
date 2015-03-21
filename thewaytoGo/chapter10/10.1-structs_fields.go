// 10.1-structs_fields.go
package main

import (
	"fmt"
)

type struct1 struct {
	i1 int
	f1 float64
	s1 string
}

func main() {
	ms := new(struct1)

	ms.i1 = 1
	ms.f1 = 1.1
	ms.s1 = "Test"

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The flot is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.s1)
	fmt.Println(ms)

}
