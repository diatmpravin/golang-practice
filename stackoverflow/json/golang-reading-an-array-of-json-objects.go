package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	type myjson struct {
		myobjects []struct {
			data map[string]string
		}
	}
	file, e := ioutil.ReadFile("dat_one_extract.json")
	if e != nil {
		fmt.Printf("File Error: [%v]\n", e)
		os.Exit(1)
	}

	dec := json.NewDecoder(bytes.NewReader(file))
	var d myjson
	dec.Decode(&d)

}
