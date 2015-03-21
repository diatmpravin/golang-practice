// Write a regexp that returns the substring for string that begins with dot and until first space.

package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("\\..* ")

	fmt.Printf("1. " + re.FindString(".d 1000=11,12") + "\n") // Must return d
	fmt.Printf("2. " + re.FindString("e 2000=11") + "\n")     // Must return nothing or ""
	fmt.Printf("3. " + re.FindString(".e2000= 11") + "\n")    // Must return nothing or ""
}
