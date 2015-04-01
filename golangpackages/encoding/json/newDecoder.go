package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func main() {
	var jsonDataStream = `
         {"Name":"Adam","Age":36,"Job":"CEO"}
         {"Name":"Eve","Age":34,"Job":"CFO"}
         {"Name":"Mike","Age":38,"Job":"COO"}
         `

	type Employee struct {
		Name string
		Age  int
		Job  string
	}

	decoder := json.NewDecoder(strings.NewReader(jsonDataStream))

	for {
		var worker Employee
		if err := decoder.Decode(&worker); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%s | %d | %s\n", worker.Name, worker.Age, worker.Job)
	}

}
