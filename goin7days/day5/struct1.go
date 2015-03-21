package main

import (
	"fmt"
)

type Person struct {
	Name  string
	Phone int64
}

func main() {
	var p Person
	//p := new(Person)
	p.Name = "GopherConIndia"
	p.Phone = 9999999999

	//p := Person{Name: "pravin", Phone: 8987898767}

	//p := Person{Phone: 8978988987}

	//p := Person{"Ankit", 7878787878}

	//p := Person{ "", 3434343434}

	fmt.Println("My name is:", p.Name)
	fmt.Println("My contact number is:", p.Phone)
}
