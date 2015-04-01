package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Employee struct {
	Name string
	Age  int
	Job  string
}

func main() {

	worker := Employee{
		Name: "Adam",
		Age:  20,
		Job:  "CEO",
	}

	encoder := json.NewEncoder(os.Stdout) //output to screen

	if err := encoder.Encode(worker); err != nil {
		fmt.Println(err)
	}

}
