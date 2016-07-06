package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"io"
)

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

	var count int

	for i := 0; i < nk[0]-1; i++ {
		for j := i + 1; j < nk[0]; j++ {
			if (intS[i]+intS[j])%nk[1] == 0 {
				count++
			}
		}
	}

	fmt.Println(count)

}
