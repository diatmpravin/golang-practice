package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	type Road struct {
		Name   string
		Number int
	}
	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}

	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.HTMLEscape(&out, b)
	fmt.Printf("%+v\n", roads)
	out.WriteTo(os.Stdout)
	fmt.Println("")
}
