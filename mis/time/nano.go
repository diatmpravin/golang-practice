package main

import (
	"fmt"
	"time"
)

func main() {
	tm := time.Now().UTC()
	t := tm.Format(time.RFC3339)
	fmt.Println(t)
}
