// 6.3-multiple_return.go
package main

import (
	"fmt"
)

var num int = 10
var num2X, num3X int

func main() {
	num2X, num3X = get2Xand3X(num)
	printValue()
	num2X, num3X = get2Xand3X_1(num)
	printValue()
}

func printValue() {
	fmt.Printf("num == %d, num2X == %d, num3X == %d \n", num, num2X, num3X)
}

func get2Xand3X(i int) (int, int) {
	return i * 2, i * 3
}

func get2Xand3X_1(i int) (x2 int, x3 int) {
	x2 = i * 2
	x3 = i * 3
	return
}
