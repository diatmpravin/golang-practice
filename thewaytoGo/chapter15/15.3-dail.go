package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:5000") // tcp ipv4
	checkConnection(conn, err)

	conn, err = net.Dial("udp", "localhost:5000") // udp
	checkConnection(conn, err)

	conn, err = net.Dial("tcp", "[2620:0:2d0:200::10]:80") // tcp ipv6
	checkConnection(conn, err)	
}

func checkConnection(conn net.Conn, err error) {
	if err != nil {
		fmt.Println("error %v connecting!")
		os.Exit(1)
	}
	fmt.Println("connection is made with %v", conn)
}

