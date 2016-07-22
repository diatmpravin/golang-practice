package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModifyString(t *testing.T) {
	var modifyStringTest = []struct {
		input, output string
	}{
		{"", ""},
		{"a", "A"},
		{"ab", "AB"},
		{"abcdE", "ABCDe"},
	}

	for _, v := range modifyStringTest {
		result := modifyString(v.input)
		assert.Equal(t, v.output, result)
	}
}
