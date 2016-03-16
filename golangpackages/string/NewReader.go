 package main

 import (
 	"fmt"
 	"strings"
 )

 func main() {
 	str := "HTTP/1.1 204 No Content\r\n\r\n"

 	reader := strings.NewReader(str)

 	l := reader.Len()
 	n, err := reader.Read([]byte("HTTP"))

 	if err != nil {
 		fmt.Println(err)
 	}

 	fmt.Println("Reader length is : ", l)
 	fmt.Println("Bytes read : ", n)

 }
