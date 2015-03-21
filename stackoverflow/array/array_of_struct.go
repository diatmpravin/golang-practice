package main

import (
	"fmt"
)

type opt struct {
	shortnm      byte
	longnm, help string
	needArg      bool
}

func main() {
	basename := []opt{
		{
			shortnm: 'a',
			longnm:  "multiple",
			needArg: false,
			help:    "Usage for a",
		},
		{
			shortnm: 'b',
			longnm:  "b-option",
			needArg: false,
			help:    "Usage for b",
		},
	}

	fmt.Println(basename)
}
