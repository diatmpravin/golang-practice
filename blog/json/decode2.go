package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Record struct {
	AuthorRaw json.RawMessage `json:"author"`
	Title     string          `json:"title"`
	URL       string          `json:"url"`

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

	// todo
	for _, v := range *x {
		var s string
    	err = json.Unmarshal(v.AuthorRaw, &s); err == nil {
        	v.AuthorEmail = s
    	}

    	var n uint64
    	err = json.Unmarshal(v.AuthorRaw, &n); err == nil {
        	v.AuthorID = n
    	}
	}

	fmt.Printf("%+v\n", *x)
}
