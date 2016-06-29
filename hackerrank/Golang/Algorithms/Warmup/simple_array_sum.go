package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"

)

// func readInt(n uint32) (sum uint32) {
// 	var temp uint32
// 	for i := 0; i < int(n); i++ {
// 		fmt.Scan(&temp)
// 		sum += temp
// 	}
// 	return
// }

func readInt(n uint32) (sum uint32) {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	for i := 0; i < int(n); i++ {
		s.Scan()
		v, _ := strconv.Atoi(s.Text())
		sum += uint32(v)
	}
	return
}

func main() {
	var n uint32
	fmt.Scan(&n)

	fmt.Println(readInt(n))
}
