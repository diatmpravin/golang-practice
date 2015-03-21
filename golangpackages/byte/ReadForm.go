 package main

 import (
    "bytes"
    "fmt"
 )

 func main() {

   bufferA := bytes.NewBuffer([]byte("abcd"))

   bufferB := bytes.NewBuffer(nil)

   n, err := bufferB.ReadFrom(bufferA)

   fmt.Printf("%v %s %d\n", err, string(bufferB.Bytes()), n)

 }