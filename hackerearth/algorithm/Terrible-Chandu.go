package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func reverseString(str string) string {
	if len(str) == 0 || len(str) == 1 {
		return str
	}

	var r bytes.Buffer
	for i := len(str) - 1; i >= 0; i-- {
		r.WriteByte(str[i])
	}

	return r.String()
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	for i := 0; i < t; i++ {
		s.Scan()
		fmt.Println(reverseString(s.Text()))
	}
}
