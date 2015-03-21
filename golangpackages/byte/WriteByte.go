 package main

 import (
   "bytes"
   "fmt"
 )

 func main() {
   buff := bytes.NewBuffer(nil) // create empty buffer

   buff.WriteByte('a')
   buff.WriteRune('中')
   buff.WriteByte('b')

   fmt.Printf("%s \n", string(buff.Bytes()))
 }