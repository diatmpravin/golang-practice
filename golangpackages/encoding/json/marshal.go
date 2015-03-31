package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type ColorGroup struct {
		ID            int      `json:",omitempty"`
		Name          string   `json:"name"`
		Colors        []string `json:"colors"`
		State         bool
		AngleBrackets string
		Ampersand     string
		BytesTest     []byte
	}

	group := ColorGroup{
		ID:            1,
		Name:          "Reds",
		Colors:        []string{"Crimson", "Red", "Ruby", "Maroon"},
		State:         true,
		AngleBrackets: "<test>",
		Ampersand:     "I & you",
		BytesTest:     'pravin',
	}

	b, err := json.Marshal(group)

	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(string(b))
}
