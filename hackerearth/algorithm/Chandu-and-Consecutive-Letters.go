package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func consecutiveLetter(str string) string {
	if len(str) == 0 || len(str) == 1 {
		return str
	}

	sliceByte := []byte{}

	for i := 0; i < len(str)-1; i++ {
		if str[i] != str[i+1] {
			sliceByte = append(sliceByte, str[i])
		}
	}

	sliceByte = append(sliceByte, str[len(str)-1])
	return string(sliceByte[:len(sliceByte)])
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	for i := 0; i < t; i++ {
		s.Scan()
		fmt.Println(consecutiveLetter(s.Text()))
	}
}
