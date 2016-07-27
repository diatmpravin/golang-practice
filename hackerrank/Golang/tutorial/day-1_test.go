package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var myPrintTest = []struct {
	input, output string
}{
	{"", ""},
	{"a", "a"},
	{"Hello World", "Hello World"},
}

func TestMyPrint(t *testing.T) {
	for _, v := range myPrintTest {
		r := myPrint(v.input)
		assert.Equal(t, v.output, r)
	}
}
