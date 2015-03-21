package main

 import (
   "crypto/aes"
   "fmt"
 )

 func main() {
    key := "opensesame45A" // not in 16, 24 or 32 bytes. Will generate error because key only have 13 bytes

    _, err := aes.NewCipher([]byte(key))

    if err != nil {
      fmt.Printf("%s\n", err.Error()) // KeySizeError function will return error message in string format
    }
 }s
