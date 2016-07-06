package main

import (
	"fmt"
)

func main() {
	var x1, v1, x2, v2 int
	fmt.Scanf("%v %v %v %v", &x1, &v1, &x2, &v2)

	j := (x1-x2) % (v2-v1)

	if x2 > x1 && v2 >= v1 {
		fmt.Println("NO")
	} else if j == 0 {
	   fmt.Println("YES")
	} else {
	   fmt.Println("NO")
	}

	// for i,j := x1, x2; ; i, j = i + v1, j + v2 {
	// 	if x2 > x1 && v2 >= v1 {
	// 		fmt.Println("NO")
	// 		break
	// 	}

	// 	if i == j {
	// 		break
	// 	}
	// }
}
