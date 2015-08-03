package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Employee struct {
	Name string
	Age  int
	Job  string
}

func main() {
	var jsonDataStream = ``

	fmt.Println("Is jsonDataStream is empty?", jsonDataStream == "")

	var worker Employee

	if err := json.NewDecoder(strings.NewReader(jsonDataStream)).Decode(&worker); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", worker)

}
