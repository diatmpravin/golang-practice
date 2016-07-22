package main

import (
	"bytes"
	"strings"
	"bufio"
	"os"
	"fmt"
)

func modifyString(str string) (r string) {
	if len(str) == 0 {
		return str
	}

	var b bytes.Buffer

	for _, v := range str {
		if v >= 'a' && v <= 'z' {
			b.WriteString(strings.ToUpper(string(v)))
		} else {
			b.WriteString(strings.ToLower(string(v)))
		}
	}

	return b.String()
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	fmt.Println(modifyString(s.Text()))
}
