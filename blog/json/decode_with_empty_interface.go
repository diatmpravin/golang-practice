package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	b := []byte(`{"Name":"Bob","Food":"Pickle"}`)
	var m interface{}

	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", m)
	fmt.Println(reflect.TypeOf(m))
}
