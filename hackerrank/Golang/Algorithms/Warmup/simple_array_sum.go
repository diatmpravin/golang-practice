package main

import (
	"bufio"
	"fmt"
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

func arraySum(s []int) (sum int) {
	for _, v := range s {
		sum = sum + v
	}

	return
}

func readInt(n uint32) (sum uint32) {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	intSlice := []int{}
	for i := 0; i < int(n); i++ {
		s.Scan()
		v, _ := strconv.Atoi(s.Text())
		intSlice = append(intSlice, v)
	}

	fmt.Println(arraySum(intSlice))
	return
}

func main() {
	var n uint32
	fmt.Scan(&n)

	readInt(n)
}
