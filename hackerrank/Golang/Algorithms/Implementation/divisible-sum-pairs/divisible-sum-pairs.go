package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func divisibleSumPairs(s []int, k int) (count int) {
	if len(s) == 0 || k == 0 {
		return 0
	}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if (s[i]+s[j])%k == 0 {
				count++
			}
		}
	}
	return
}

func convertToSlice(s []string) []int {
	arr := make([]int, len(s))
	for i, v := range s {
		r, err := strconv.ParseInt(strings.TrimSpace(v), 10, 32)
		if err != nil {
			panic(err)
		}
		arr[i] = int(r)
	}
	return arr
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	nk := convertToSlice(strings.Fields(s))

	arrLine, err := reader.ReadString('\n')
	if err != io.EOF && err != nil {
		panic(err)
	}

	intS := convertToSlice(strings.Fields(arrLine))

	fmt.Println(divisibleSumPairs(intS, nk[1]))

}
