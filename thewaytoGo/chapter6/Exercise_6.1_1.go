// Exercise_6.1_1.go
package main

import (
	"fmt"
)

func main() {
	sum, pro, diff := intOperation(5, 2)
	fmt.Println(sum, pro, diff)
}

func intOperation(x, y int) (sum, pro, diff int) {
	sum = x + y
	pro = x * y
	diff = x - y
	return
}
