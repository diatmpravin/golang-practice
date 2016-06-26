package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter somthing")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	fmt.Println(scan.Text())
}