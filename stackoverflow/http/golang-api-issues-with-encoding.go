package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	UserID int
	ID     int
	Title  string
	Body   string
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
    fmt.Println("---> Header: %#v\n", r.Header)
	fmt.Printf("---> Body: %#v\n", r.Body)
	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	post := new(Post) // or &Foo{}
	err := getJson("http://jsonplaceholder.typicode.com/posts/1", &post)
	if err != nil {
        panic(err)
    }

    fmt.Printf("Post: %+v\n", post)
}
