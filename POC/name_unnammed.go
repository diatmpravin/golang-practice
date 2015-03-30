package main

import "fmt"

type Map map[string]string

//this is valid
// func (m Map) Display(s string) {
// 	fmt.Println(s)
// }

// invalide
func (m map[string]string) Display(s string) {
  fmt.Println(s)
}

func main() {
	var m = Map{"first_name": "Rob"}

	m.Display("Hello")
}
