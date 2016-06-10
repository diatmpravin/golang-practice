package factorial

import (
	"fmt"
)

func firstFactorial(num int) int{
	var fact int = 1

	if num == 0 {
		return 1
	} else {
		for i := 1; i <= num ; i++ {
			fact = fact * i	
		}
	}

	fmt.Printf("Fact of %d is %d \n", num, fact)
	return fact
}