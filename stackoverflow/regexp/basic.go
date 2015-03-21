package main

import (
	"fmt"
	"regexp"
)

func main() {
	rp := regexp.MustCompile("[a-z]+")   // Panics if the regex is incorrect
	rp1, err := regexp.Compile("[a-z]+") // Returns an error message if the regex is incorrect
	if err != nil {
		fmt.Println(err)
	}

	foundBool := rp.MatchString("abc") // Report weather regex matches the string
	foundBool1 := rp1.MatchString("abc12")

	fmt.Println(foundBool, foundBool1)

	fmt.Println(rp.FindString("123 abc def"))           // Find the first string in text
	fmt.Println(rp.FindAllString("123 abc 43 def", -1)) // Return slice of(all) string from text

	rp2 := regexp.MustCompile("([a-z])([a-z])[a-z]+")
	fmt.Println(rp2.FindAllStringSubmatch("abc", -1)) // ["abc", "a", "b"]

	rp3 := regexp.MustCompile("([a-z])([a-z])[a-z]+")
	fmt.Println(rp3.FindAllStringSubmatchIndex("abc", -1)) // [0, 3, 0, 1, 1, 2]
}
