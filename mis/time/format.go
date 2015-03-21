package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func main() {
	tm := time.Now().UTC()
	fmt.Println(tm)
	t := tm.Format(time.StampMilli)
	fmt.Println(t)
	r := strings.Split(t, " ")
	fmt.Println(r)
	fmt.Println(reflect.TypeOf(t))
}
