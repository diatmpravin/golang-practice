package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsondata = []byte(`[
         {"Name":"Adam","Age":36,"Job":"CEO"},
         {"Name":"Eve","Age":34,"Job":"CFO"},
         {"Name":"Mike","Age":38,"Job":"COO"}
         ]`)

	type Employee struct {
		Name string
		Age  int
		Job  string
	}

	var workers []Employee

	err := json.Unmarshal(jsondata, &workers)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v \n", workers)
}
