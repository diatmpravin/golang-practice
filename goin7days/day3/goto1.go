package main

import "fmt"

func main() {
	j := 0
GOLANG:
	j = j + 1
	for i := 0; i < 5; i++ {
		if j > 4 {
			break
		}

		if i == 2 {
			goto GOLANG
		}
		fmt.Println("Value of I is:", i)
	}
}
