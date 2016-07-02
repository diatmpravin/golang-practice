package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanLines)
	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	if t >= 1 && t <= 100 {
		for i := 0; i < t; i++ {
			s.Scan()
			str := s.Text()

			if utf8.RuneCountInString(str)%2 == 0 {
				a := str[0 : len(str)/2]
				b := str[len(str)/2:]

				ar := map[rune]int{}
				br := map[rune]int{}


				for _, r := range a {
					ar[r]++
				}

				for _, r := range b {
					br[r]++
				}

				tot := 0
				for k, v := range ar {
					bc := br[k]
					if v > bc {
						tot += v - bc
					}
				}
				fmt.Println(tot)
			} else {
				fmt.Println(-1)
			}

		}
	}
}