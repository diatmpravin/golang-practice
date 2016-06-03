package reverse

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFirstReverse(t *testing.T){
	var reverseTest = []struct{
		in, out string
	}{
		{"Hello", "olleH"},
		{"世界", "界世"},
		{"", ""},
	}

	for _, v := range reverseTest {
		result := firstReverse(v.in)	
		assert.Equal(t, result, v.out)
	}	
}