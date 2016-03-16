package main

import (
	"fmt"
	"time"
)

var x = [5]int{11, 32, 43, 45}

func main() {
	for i := range x {
		go func(){
			fmt.Println(i)
		}()
	}
	time.Sleep(5 * time.Second)
}
