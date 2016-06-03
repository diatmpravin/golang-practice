package factorial

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFirstFactorial(t *testing.T) {
	var factorialTest = []struct{
		in, out int
	}{
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	}

	for _, v := range factorialTest {
		result := firstFactorial(v.in)
		if v.out != result {
			assert.Equal(t, v.in, result)
			fmt.Println("Expected factorial of %d is %d but got--> %d", v.in, v.out, result)
		}
	}
}

