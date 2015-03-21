// 6.17-return_defer.go
package main

import (
	"fmt"
)

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func main() {
	fmt.Println(f())
}
