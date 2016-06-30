package main

import (
	"fmt"
	"time"
)

func main() {
	var a string

	fmt.Scanf("%v\n", &a)

	ft, err := time.Parse("03:04:05PM", a)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ft.Format("15:04:05"))
}
