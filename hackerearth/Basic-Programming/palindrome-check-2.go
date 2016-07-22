package main

import (
	"math"
	"bufio"
	"os"
	"fmt"
)

func palindrome(str string) (r string) {
	if len(str) == 0 {
		return "NO"
	}

	if len(str) == 1 {
		return "NO"
	}

	var start int 

	for i := len(str)-1; i >= int(math.Abs(float64(len(str)/2))); i-- {
		if str[i] != str[start] {
			return "NO"
		}
		start++	
	}

	return "YES"
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()

	fmt.Println(palindrome(s.Text())) 
}