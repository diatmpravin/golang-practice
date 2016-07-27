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
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(myPrint(text))
}
