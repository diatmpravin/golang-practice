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
		Age:  36,
		Job:  "CEO",
	}

	output, err := json.MarshalIndent(worker, "%%%%", "$$") // <--- here
	// output, err := json.Marshal(worker)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}
