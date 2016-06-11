package main

import (
	"fmt"
	"github.com/buger/jsonparser"
)

var data = []byte(`{
	"object": {
		"buffer_size": 10,
		"Databases": [{
			"host": "localhost",
			"user": "root",
			"pass": "",
			"type": "mysql",
			"name": "go",
			"Tables": [{
				"name": "testing",
				"statment": "teststring",
				"regex": "teststring ([0-9]+) ([A-z]+)",
				"Types": [{
					"id": "int",
					"value": "string"
				}]
			}]
		}]
	}
}`)

func main() {
	val, _, _, err := jsonparser.Get(data, "object", "Databases")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(val))
}
