package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertToInt(s []string) []int {
	arr := []int{}
	for _, val := range s {
		v, _ := strconv.Atoi(val)
		arr = append(arr, v)
	}

	return arr
}

func vehicleCheck(i, j int, arr []int) (v int) {
	v = arr[i]
	
	arr1 := arr[i:j+1]

	for _, val := range arr1 {
		if v >= val {
			v = val
		}
	}
	return
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanLines)

	s.Scan()
	nt := convertToInt(strings.Split(s.Text(), " "))

	s.Scan()
	width := convertToInt(strings.Split(s.Text(), " "))

	for i := 0; i < nt[1]; i++ {
		s.Scan()
		ij := convertToInt(strings.Split(s.Text(), " "))

		if ij[0] >= 0 && ij[1] > ij[0] && ij[1] < nt[0] {
			fmt.Println(vehicleCheck(ij[0], ij[1], width))
		}
	}
}
