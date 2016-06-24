package main 

import (
	"fmt"
	"unicode/utf8"
	"bytes"
)

func main() {
	search_input := []byte("abcdefghidefdef")
    fmt.Println("---->", search_input)

    key_slice := []byte("def")
    fmt.Println("---->", key_slice)

    fmt.Println(count(search_input,key_slice))
    fmt.Println(bytes.IndexByte(search_input,key_slice[0]))
}

// Count counts the number of non-overlapping instances of sep in s.
// If sep is an empty slice, Count returns 1 + the number of Unicode code points in s.
func count(s, sep []byte) int {
	n := len(sep)
	fmt.Println("Length of key_slice is:", n)
	fmt.Println("Length of seach_inp is:", len(s))
	if n == 0 {
		return utf8.RuneCount(s) + 1
	}
	if n > len(s) {
		return 0
	}
	count := 0
	c := sep[0]
	i := 0
	t := s[:len(s)-n+1]
	fmt.Println("c----->", c)
	fmt.Println("t----->", t)
	for i < len(t) {
		fmt.Println("************** LOOP **************", i)
		if t[i] != c {
			o := bytes.IndexByte(t[i:], c)
			fmt.Println("************** INDEX **************", o)
			if o < 0 {
				break
			}
			i += o
		}
		fmt.Println("************** STRING ************** and i====>", s[i:i+n], i)
		if n == 1 || bytes.Equal(s[i:i+n], sep) {
			count++
			i += n
			continue
		}
		i++
	}
	return count
}