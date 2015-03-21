package main

 import (
   "bytes"
   "fmt"
 )

 func main() {
   buff := bytes.NewBufferString("abcdefghi")

   unread := buff.Bytes()

   fmt.Println(string(unread)) // since nothing has been read. it will return everything

   var readbuff [3]byte

   // buff.Read(readbuff[0:2])

   fmt.Println(readbuff)

   unread2 := buff.Bytes() // this time only return the unread portion

   fmt.Println(string(unread2))
 }