package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(b)
	// byte alias for uint8
	fmt.Println(reflect.TypeOf(b))
}
