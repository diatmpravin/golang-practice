package main

import (
	"encoding/json"
	"fmt"
)

type App struct {
	Id    string `json:"id"`
}

type Org struct {
	Title string `json:"title"`
}

type AppWithOrg struct {
	App
	Org
}

func main() {
	data := []byte(`{"id": "k34rAT4", "title": "My Awesome App"}`)

	var appwithorg AppWithOrg

	err := json.Unmarshal(data, &appwithorg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", appwithorg)

	
}
