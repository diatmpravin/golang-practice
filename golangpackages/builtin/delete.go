package main

import "fmt"

type Person struct {
	Age int
}

var personmap map[string]Person

func main() {
	personmap = make(map[string]Person)

	personmap["Adam"] = Person{36}

	fmt.Println("Before delete : ", personmap)

	delete(personmap, "Adam")

	fmt.Println("After delete : ", personmap)

}
