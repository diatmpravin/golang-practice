package main

import (
	"encoding/json"
	"fmt"
)

type App struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

func main() {
	data := []byte(`{"id": "k34rAT4", "title": "My Awesome App"}`)

	var app App

	err := json.Unmarshal(data, &app)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", app)

	a, err := json.Marshal(app)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", string(a))
}
