package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {

	// {"Name":"Bob","Food":"Pickle"}
	// {"Name":"Alice","Body":"Hello","Time":1294706395881547000}
	// {"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}

	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			if k != "Name" {
				delete(v, k)
			}
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}
