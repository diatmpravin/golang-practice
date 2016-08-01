package main

import (
	"fmt"
)

func willTheyMeet(x1, v1, x2, v2 int) string {
	if x2 > x1 && v2 >= v1 || x1 > x2 && v1 >= v2 {
		return "NO"
	}

	for i, j := x1, x2; ; i, j = i+v1, j+v2 {

		if j > 10001 {
			break
		}

		if i == j {
			return "YES"
		}
	}

	return "NO"
}

func main() {
	var x1, v1, x2, v2 int
	fmt.Scanf("%v %v %v %v", &x1, &v1, &x2, &v2)
	fmt.Println(willTheyMeet(x1, v1, x2, v2))
}
