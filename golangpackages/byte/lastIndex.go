package main

 import (
         "bytes"
         "fmt"
 )

 func main() {

         str := []byte("abcdefg...xyz")
         fmt.Println(lastIndex(str, []byte("x")))

         str2 := []byte("abcxyz")
         fmt.Println(lastIndex(str2, []byte("x")))
 }

// lastIndex returns the index of the last instance of sep in s, or -1 if sep is not present in s.
func lastIndex(s, sep []byte) int {
	n := len(sep)
	if n == 0 {
		return len(s)
	}
	c := sep[0]
	for i := len(s) - n; i >= 0; i-- {
		if s[i] == c && (n == 1 || bytes.Equal(s[i:i+n], sep)) {
			return i
		}
	}
	return -1
}