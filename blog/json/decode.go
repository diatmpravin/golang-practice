package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Record struct {
	Author string
	Title  string
	URL    string
}

func main() {
	a := `{"author": "attila@attilaolah.eu", "title": "My Blog", "url": "http://attilaolah.eu"}`

	x := new(Record)
	err := json.NewDecoder(strings.NewReader(a)).Decode(x)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", *x)
}
