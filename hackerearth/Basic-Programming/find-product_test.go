package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindProduct(t *testing.T) {
	var findProductTest = []struct {
		input  []int
		output int
	}{
		{[]int{}, 0},
		{[]int{2}, 2},
		{[]int{9}, 9},
		{[]int{1, 2}, 2},
		{[]int{1, 2, 3, 4, 5}, 120},
	}

	for _, v := range findProductTest {
		result := findProduct(v.input)
		assert.Equal(t, v.output, result)
	}
}
