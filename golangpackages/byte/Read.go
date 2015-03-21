package main

import (
	"bytes"
	"fmt"
)

func main() {
	buff := bytes.NewBufferString("abcdefgh")

	var threebytes [3]byte

	n, err := buff.Read(threebytes[0:3])
	fmt.Printf("%v  %s\n", err, string(threebytes[:n]))

	fmt.Println(buff)

	n, err = buff.Read(threebytes[:])
	fmt.Printf("%v  %s\n", err, string(threebytes[:n]))

	n, err = buff.Read(threebytes[:])
	fmt.Printf("%v  %s\n", err, string(threebytes[:n]))

}
