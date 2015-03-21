// 6.8-defer.go
package main

import "fmt"

func main() {

	Function1()
}
func Function1() {
	fmt.Println("In Function1 at the top")
	defer Function2()
	fmt.Println("In Function1 at the bottom!")
}
func Function2() {

	fmt.Println("Function2: Deferred until the end of the calling function!")
}
