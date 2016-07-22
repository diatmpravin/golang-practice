package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindrome(t *testing.T) {
	var palindromeTest = []struct {
		input, output string
	}{
		{"", "NO"},
		{"a", "NO"},
		{"anna", "YES"},
		{"ancna", "YES"},
		{"kayak", "YES"},
		{"mom", "YES"},
		{"racecar", "YES"},
		{"I did did I", "YES"},
		{"step on no pets", "YES"},
	}

	for _, v := range palindromeTest {
		result := palindrome(v.input)
		assert.Equal(t, v.output, result)
	}
}
