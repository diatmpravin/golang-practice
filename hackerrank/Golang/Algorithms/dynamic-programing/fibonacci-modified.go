package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math/big"
)

var s *bufio.Scanner

func convertToInt(s []string, n int64) []int64 {
	arr := make([]int64, n)
	var i int64

	for i = 0; i < n; i++ {
		val, _ := strconv.Atoi(s[i])
		arr[i] = int64(val)
	}
	return arr
}

func calculateFab(a, b, n int64) (temp *big.Int) {
	var i int64

    tn  := big.NewInt(a)
    tn1 := big.NewInt(b)

    for i = 2; i < n; i += 1 {
    	temp = new(big.Int)

        temp.Mul(tn1, tn1)
        temp.Add(temp, tn)

        tn = tn1
        tn1 = temp
    }

    return temp
}

func main() {
	s = bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanLines)

	s.Scan()
	abn := convertToInt(strings.Split(s.Text(), " "), int64(3))

	fmt.Println(calculateFab(abn[0], abn[1], abn[2]))	
}
