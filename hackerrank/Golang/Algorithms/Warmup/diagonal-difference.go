package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum(n int, arr [][]int) int {
	var primeDiagonal, secondryDiagonal int

	for i := 0; i < n; i++ {
		primeDiagonal += arr[i][i]
	}


	for i, j := (n-1), 0; i >= 0; i, j = i-1, j+1 {
		secondryDiagonal += arr[j][i]
	}


	return primeDiagonal - secondryDiagonal
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([][]int, n)

	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			s.Scan()
			v, _ := strconv.Atoi(s.Text())
			arr[i][j] = v			 			
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(arr[i][j])			 			
		}
		fmt.Println()
	}

	fmt.Println(sum(n, arr))
}
