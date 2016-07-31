package main

import (
	"bufio"
	"fmt"
	// "math/big"
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

		var newN int64

		if l == r {
			tmp := sliceInt[l-1 : r]
			tmp1, _ := strconv.Atoi(tmp[0])
			newN = int64(tmp1)
		} else {
			tmp := strings.Join(sliceInt[l-1:r], "")
			tmp1, _ := strconv.Atoi(tmp)
			newN = int64(tmp1)
		}

		if isDivisibleBy7(newN) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func isDivisibleBy7(num int64) bool {
	// if num < 0 {
	// 	return isDivisibleBy7(-num)
	// }

	// if num == 0 || num == 7 {
	// 	return true
	// }

	// if num < 10 {
	// 	return false
	// }

	// return isDivisibleBy7(num/10 - 2*(num-num/10*10))

	num = (num&7) + ((num>>3)&7) + (num>>6)
    num = (num&7) + (num>>3)

    if num == 7 {
    	return true
    } else {
    	return false
    }
}
