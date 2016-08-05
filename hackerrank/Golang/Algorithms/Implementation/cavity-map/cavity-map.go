package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	Map := make([][]rune, n)
	MapCopy := make([][]rune, n)

	var row string

	for i := 0; i < n; i++ {
		Map[i] = make([]rune, n)
		MapCopy[i] = make([]rune, n)
		fmt.Scan(&row)

		for j := 0; j < n; j++ {
			Map[i][j] = rune(row[j])
			MapCopy[i][j] = rune(row[j])
		}
	}

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			gap1 := Map[i][j] - Map[i][j-1]
			gap2 := Map[i][j] - Map[i-1][j]
			gap3 := Map[i][j] - Map[i+1][j]
			gap4 := Map[i][j] - Map[i][j+1]

			if gap1 > 0 && gap2 > 0 && gap3 > 0 && gap4 > 0 {
				MapCopy[i][j] = rune('X')
			}

		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%c", MapCopy[i][j])
		}
		fmt.Println()
	}
}
