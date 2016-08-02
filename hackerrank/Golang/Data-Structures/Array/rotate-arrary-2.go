package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var s *bufio.Scanner

func convertToInt(s []string) []int {
	arr := []int{}
	for _, v := range s {
		val, _ := strconv.Atoi(v)
		arr = append(arr, val)
	}

	return arr
}

func siftArry(a, b, c []int) []int {
	for j := 0; j < a[0]; j++ {
		b[a[0]-(1+j)] = c[j]
		fmt.Println("--->", b)
	}

	return b
}

func display(a, b []int) {
	for i := 0; i < a[0]; i++ {
		fmt.Print(b[i], " ")
	}
}

func main() {
	s = bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanLines)

	s.Scan()
	nd := convertToInt(strings.Split(s.Text(), " "))

	s.Scan()
	arrElement := convertToInt(strings.Split(s.Text(), " "))

	finalArry := make([]int, nd[0])
	for i := 0; i < nd[1] ; i++ {
		newArr := make([]int, nd[0])
		a := siftArry(nd, newArr, arrElement)
		arrElement = a
		finalArry = a 
	}

	display(nd, finalArry)
}
