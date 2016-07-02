package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
	"strings"
)

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanLines)
	s.Scan()
	w := s.Text()
	newString := w

	stringSlice := []byte{}
	for i := 0; i < utf8.RuneCountInString(w)-1; i++ {
		stringSlice[i] = w[i]
	}

	for i := 0; i < len(stringSlice); i++ {
		if stringSlice[i] == stringSlice[i+1] {
			newString = strings.Replace(newString, string(w[i]), "", 1)
			newString = strings.Replace(newString, string(w[i+1]), "", 1)
		}
		fmt.Println("------>", newString)
	}

	fmt.Println(newString)
}
