package main

 import (
         "bytes"
         "fmt"
 )

 func main() {

         str := []byte("abcdefg...xyz")
         fmt.Println(bytes.LastIndex(str, []byte("x")))

         str2 := []byte("abcxyz")
         fmt.Println(bytes.LastIndex(str2, []byte("x")))
 }