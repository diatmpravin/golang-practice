package main

import "fmt"
import "regexp"

func main() {

	re := regexp.MustCompile(`\.(.*) `)
	match := re.FindStringSubmatch(".pravind 1000=11,12")
	if len(match) != 0 {
		fmt.Printf("1. %s\n", match[1])
	}

	match = re.FindStringSubmatch("e 2000=11")
	if len(match) != 0 {
		fmt.Printf("2. %s\n", match[1])
	}

	match = re.FindStringSubmatch(".e2000=11")
	if len(match) != 0 {
		fmt.Printf("3. %s\n", match[1])
	}
}
