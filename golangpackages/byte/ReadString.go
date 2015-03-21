package main

 import (
    "bytes"
    "fmt"
 )

 func main() {

   buff := bytes.NewBuffer([]byte("abcde#fg"))

   line, err := buff.ReadString('#')

   fmt.Println(line,err)

 }