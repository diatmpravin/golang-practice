package main

import (
	"fmt"
)

func readPhoneBook(n int) map[string]int {
	phoneBook := make(map[string]int)

	for i := 0; i < n; i++ {
		var key string
		var value int
		fmt.Scan(&key, &value)
		phoneBook[key] = value
	}
	return phoneBook
}

func findPhoneNo(p map[string]int) {
	for {
		var phone string
		fmt.Scan(&phone)

		_, ok := p[phone]
		if ok {
			fmt.Printf("%s=%d \n", phone, p[phone])
		} else {
			fmt.Println("Not found")
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	phoneBook := readPhoneBook(n)

	findPhoneNo(phoneBook)
}
