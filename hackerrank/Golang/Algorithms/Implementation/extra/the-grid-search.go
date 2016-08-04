package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type testCase struct {
	grid    []string
	pattern []string
}

func main() {
	var numberOfTestCases int

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	fmt.Sscanf(line, "%d", &numberOfTestCases)

	testCases := make([]string, numberOfTestCases)
	for i := 0; i < numberOfTestCases; i++ {
		var numberOfRows, numberOfColumns int

		line, _ := reader.ReadString('\n')
		fmt.Sscanf(line, "%d %d", &numberOfRows, &numberOfColumns)
		grid := readMatrix(reader, numberOfRows, numberOfColumns)

		line, _ = reader.ReadString('\n')
		fmt.Sscanf(line, "%d %d", &numberOfRows, &numberOfColumns)
		pattern := readMatrix(reader, numberOfRows, numberOfColumns)

		testCases[i] = hasPattern(grid, pattern)
	}

	for _, testCase := range testCases {
		fmt.Println(testCase)
	}
}

func readMatrix(reader *bufio.Reader, numberOfRows, numberOfColumns int) (matrix []string) {
	matrix = make([]string, numberOfRows)
	for i := 0; i < numberOfRows; i++ {
		line, _, _ := reader.ReadLine()
		matrix[i] = string(line)
	}

	return matrix
}

func readMatrixFromFile(reader *bufio.Reader, numberOfRows, numberOfColumns int) (matrix []string) {
	matrix = make([]string, numberOfRows)
	for i := 0; i < numberOfRows; i++ {
		line, _, _ := reader.ReadLine()
		matrix[i] = string(line)
	}

	return matrix
}

func hasPattern(grid, pattern []string) string {
	gridHeight := len(grid)
	patternHeight := len(pattern)

	for rowIndex := 0; rowIndex < gridHeight; rowIndex++ {
		index := strings.Index(grid[rowIndex], pattern[0])
		if index != -1 {
			has := true
			for i := 0; i < patternHeight; i++ {
				index2 := strings.Index(grid[rowIndex+i], pattern[i])
				if index2 == -1 || index2 != index {
					has = false
					break
				}
			}
			if has {
				return "YES"
			}
		}
	}

	return "NO"
}
