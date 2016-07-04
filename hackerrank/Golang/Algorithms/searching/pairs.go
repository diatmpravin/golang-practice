package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	K, _ := strconv.Atoi(scanner.Text())

	numbers := make(map[int]int)
	for i := 0; i < N && scanner.Scan(); i++ {
		v, _ := strconv.Atoi(scanner.Text())
		numbers[v] = 0
	}

	answer := 0
	for key, _ := range numbers {
		if _, ok := numbers[key-K]; ok {
			answer++
		}
	}

	fmt.Println(answer)
}
