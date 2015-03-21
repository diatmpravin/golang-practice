// Exercise_6.5.go
package main

import (
	"fmt"
)

func main() {
	printrec(10)
}

func printrec(i int) {
	for ; i > 0; i-- {
		fmt.Println(i)
	}
}
