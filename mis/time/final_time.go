package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	tm := time.Now().UTC()

	w := tm.Format(time.RFC3339)

	ws := strings.Split(w, "T")

	t := tm.Format(time.StampMilli)

	r := strings.Split(t, " ")

	s := []string{ws[0], r[2]}
	final := strings.Join(s, "T")

	fmt.Println(final + "Z")
}
