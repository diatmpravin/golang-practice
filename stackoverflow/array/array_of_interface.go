package main

import "fmt"

type user struct {
	name, email string
}

func main() {
	u := []interface{}{"pravin", "pravinmishra@gmail.com", 1}

	fmt.Println(u)
}
