package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"unicode/utf8"
)

func pangrams(s string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	trimed := strings.Replace(s, " ", "", -1)
	lowcase := strings.ToLower(trimed)

	for _, v := range lowcase {
		alphabet = strings.Replace(alphabet, string(v), "", -1)	
	}

	if utf8.RuneCountInString(alphabet) == 0 {
		return "pangram"
	} else {
		return "not pangram"
	}
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanLines)
	s.Scan()

	fmt.Println(pangrams(s.Text()))
}