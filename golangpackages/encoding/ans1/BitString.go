package main

 import (
    "encoding/asn1"
    "fmt"
 )

 func main() {
    var bstr asn1.BitString

    bstr.Bytes = []byte{0x82, 0x40}
    bstr.BitLength = 16

    fmt.Println(bstr.At(0))
    fmt.Println(bstr.RightAlign)

 }