package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	reader := strings.NewReader(`{"Name" : "Adam"} extra string to be buffered`)

	var m struct {
		Name string
	}

	decoder := json.NewDecoder(reader)
	// Reads the next JSON-encoded value from its input
	err := decoder.Decode(&m)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(m)

	// Buffered returns a reader of the data remaining in the Decoder's buffer.
	extraString, err := ioutil.ReadAll(decoder.Buffered()) // read the remainder

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(extraString))
}
