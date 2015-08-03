package main

import (
	"fmt"
	"net/http"
)

func hander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", hander)
	http.ListenAndServe(":8080", nil)
}
