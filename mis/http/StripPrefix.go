package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	fmt.Println("FS-->", fs)
	http.Handle("/public/", http.StripPrefix("/public/", fs))
}
