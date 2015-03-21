package main

 import (
    "bytes"
    "fmt"
 )

 func main() {

   buff := bytes.NewBuffer([]byte("abcdefg"))

   rune, _, _ := buff.ReadRune()

   fmt.Println(string(rune))

   // remainder runes

   fmt.Println(string(buff.Bytes()))

   buff.UnreadRune()

   // back to beginning
   fmt.Println(string(buff.Bytes()))

 }