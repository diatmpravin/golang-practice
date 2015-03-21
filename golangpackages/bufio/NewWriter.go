package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	screenWriter := bufio.NewWriter(os.Stdout)
	fmt.Fprint(screenWriter, "Hello, ")
	fmt.Fprint(screenWriter, "world! 你好世界!\n")
	screenWriter.Flush() // Don't forget to flush!
}
