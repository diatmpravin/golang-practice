package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	NAme string `json:"name"`
	Body string
	Time int64
}

func main() {
	b := []byte(`{"name":"Bob","Food":"Pickle"}`)
	var m Message

	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("%+v\n", m)
}
