// Exercise_10.1.go
package main

import (
	"fmt"
)

type Address struct {
	Name      string
	Dob       string
	ContactNo string
	City      string
	State     string
}

type VCard struct {
	Address Address
}

func main() {
	vcard := VCard{Address{"Pravin", "15/05/1988", "+91-7042940011", "Gurgaon", "Haryana"}}
	fmt.Printf("%+v\n", vcard)
}
