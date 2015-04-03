package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Record struct {
	AuthorRaw interface{} `json:"author"`
	Title     string      `json:"title"`
	URL       string      `json:"url"`

	AuthorEmail string
	AuthorId    uint64
}

func main() {
	a := `[
			{"author": "attila@attilaolah.eu", "title": "My Blog", "url": "http://attilaolah.eu"},
			{"author": 1234567890, "title":  "Westartup", "url":    "http://www.westartup.eu"}
			]`

	x := new([]Record)

	err := json.NewDecoder(strings.NewReader(a)).Decode(x)
	if err != nil {
		fmt.Println(err)
	}


	// access as address
	for _, v := range *x {
		switch t := v.AuthorRaw.(type) {
		case string:
			v.AuthorEmail = t
		case float64:
			v.AuthorId = uint64(t)
		}
	}

	fmt.Printf("%+v\n", *x)
}
