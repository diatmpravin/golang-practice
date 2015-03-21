package main

import "fmt"

func main() {
	var x [5]float64
	scores := [5]float64{98, 93, 77, 82, 83}

	for i, _ := range x {
		for j, _ := range scores {
			fmt.Printf("i=> %d and J=> %d\n",i, j)
			x[j] = scores[j]
		}
	}
	fmt.Println(x)
}
