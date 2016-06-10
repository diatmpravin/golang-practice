package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Round(time.Millisecond).Format("15:04:05.999999999"))
}
