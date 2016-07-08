package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Integer value into array of its digits
func intToArray(value int) []int {
    n := len(strconv.Itoa(value))
    
    array := make([]int, n)
    
    for i := 0; i < n; i++ {
        array[i] = value % 10
        value /= 10
    }
    
    return array
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	for i := 0; i < t; i++ {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())

		digitArray := intToArray(n)

		var count int
		for _, value := range digitArray { 
			if value != 0 &&  n%value == 0 {
				count++
			}
		}
		fmt.Println(count)
	}
}