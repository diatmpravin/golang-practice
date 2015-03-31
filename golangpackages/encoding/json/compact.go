package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	dst := new(bytes.Buffer)

	b := []byte(`{"name":"Bob", "Food":"Pickle"}`)

	err := json.Compact(dst, b)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dst)
}
