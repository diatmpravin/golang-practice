package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	sliceInt := strings.Split(s.Text(), "")

	s.Scan()
	q, _ := strconv.Atoi(s.Text())

	for i := 0; i < q; i++ {
		s.Scan()
		l, _ := strconv.Atoi(s.Text())

		s.Scan()
		r, _ := strconv.Atoi(s.Text())

		var newN *big.Int
		// var tmp1 int64

		if l == r {
			tmp := sliceInt[l-1 : r]
			tmp1, _ := strconv.Atoi(tmp[0])
			newN1 := int64(tmp1)
			// tmp1, _ := strconv.Atoi(tmp[0])
			newN = big.NewInt(newN1)
		} else {
			tmp := strings.Join(sliceInt[l-1:r], "")
			tmp1, _ := strconv.Atoi(tmp)
			newN1 := int64(tmp1)
			// tmp1, _ := strconv.Atoi(tmp)
			newN = big.NewInt(newN1)
		}

		if isDivisibleBy7(newN) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func isDivisibleBy7(num *big.Int) bool {
	if num < 0 {
		return isDivisibleBy7(-num)
	}

	if num == 0 || num == 7 {
		return true
	}

	if num < 10 {
		return false
	}

	return isDivisibleBy7(num/10 - 2*(num-num/10*10))
}
