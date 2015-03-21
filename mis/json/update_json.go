package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type Color struct {
		Space string
		Point json.RawMessage // delay parsing until we know the color space
	}
	type RGB struct {
		R uint8
		G uint8
		B uint8
	}
	type YCbCr struct {
		Y  uint8
		Cb int8
		Cr int8
	}

	var j = []byte(`
        {"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}}`)
	var colors Color
	err := json.Unmarshal(j, &colors)
	if err != nil {
		log.Fatalln("error:", err)
	}
	colors.Space = "no-space"

	b, err := json.Marshal(&colors)
	if err != nil {
		panic(err)
	}
	fmt.Printf("b is now %s", b)
	return

}
