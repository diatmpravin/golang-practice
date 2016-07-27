package main

import (
	"bufio"
	"fmt"
	"os"
)

func myPrint(s string) string {
	return s
}

func main() {
	fmt.Println("Hello, World.")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	fmt.Println(myPrint(scan.Text()))
}