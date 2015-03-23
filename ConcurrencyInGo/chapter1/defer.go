package main

import(
	"fmt"
)

func main(){
	aValue := new(int)

        defer fmt.Println("Final Value:", *aValue)

	for i := 0 ; i <100 ; i++ {
		fmt.Println(*aValue)
		*aValue++
	}
}
