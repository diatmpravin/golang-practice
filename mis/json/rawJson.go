package main

import (
	"encoding/json"
	"fmt"
)

const data = `{"object": 
    {
       "buffer_size": 10,
       "Databases":
       [
               {
                       "host": "localhost",
                       "user": "root",
                       "pass": "",
                       "type": "mysql",
                       "name": "go",
                       "Tables":
                       [
                               {
                                       "name": "testing",
                                       "statment": "teststring",
                                       "regex": "teststring ([0-9]+) ([A-z]+)",
                                       "Types": 
                                        [
                                           {
                                               "id": "int",
                                               "value": "string"
                                           }
                                        ]
                               }
                       ]
               }
       ]
    }
}`

func main() {
	var objmap map[string]json.RawMessage
	err := json.Unmarshal([]byte(data), &objmap)

	fmt.Println(err)
	fmt.Println(string(objmap["object"]))
}
