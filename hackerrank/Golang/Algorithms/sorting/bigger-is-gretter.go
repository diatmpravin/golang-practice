package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	s.Split(bufio.ScanLines)

	for i := 0; i < n; i++ {
		s.Scan()
		word := s.Text()

		cSlice := make([]byte, len(word))
		for j := 0; j < len(word); j++ {
			cSlice[j] = word[j]
		}

		found := false
		for j := len(cSlice) - 2; j > -1; j-- {
			k := j
			v := cSlice[k]
			if v >= cSlice[len(cSlice)-1] {
				for ; k < len(cSlice)-1; k++ {
					cSlice[k] = cSlice[k+1]
				}
				cSlice[len(cSlice)-1] = v
			} else {
				for k++; k < len(cSlice) && v >= cSlice[k]; k++ {
				}
				cSlice[j], cSlice[k] = cSlice[k], v
				found = string(cSlice) != word
				break
			}
		}
		if found {
			fmt.Printf("%v\n", string(cSlice))
		} else {
			fmt.Printf("no answer\n")
		}
	}
}
