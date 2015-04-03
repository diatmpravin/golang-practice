package main

import (
	"encoding/json"
	"fmt"
)

type PublicKey struct {
	Id  int
	Key string
}

func main() {
	keysBody := []byte(`[{"id": 1,"key": "-"},{"id": 2,"key": "-"},{"id": 3,"key": "-"}]`)
	// slice of struct
	keys := make([]PublicKey, 0)
	json.Unmarshal(keysBody, &keys)
	fmt.Printf("%+v\n", keys)
}
