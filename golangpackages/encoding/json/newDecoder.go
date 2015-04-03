package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"reflect"
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

	fmt.Println(reflect.TypeOf(jsonDataStream))

	a := strings.NewReader(jsonDataStream)
	fmt.Println(reflect.TypeOf(a))

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
