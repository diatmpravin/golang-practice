package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivisibleSumPairs(t *testing.T) {
	var divisibleSumPairsTest = []struct {
		s         []int
		k, output int
	}{
		{[]int{}, 3, 0},
		{[]int{}, 0, 0},
		{[]int{1, 3, 2}, 0, 0},
		{[]int{1, 3, 2, 6, 1, 2}, 3, 5},
	}

	for _, v := range divisibleSumPairsTest {
		r := divisibleSumPairs(v.s, v.k)
		assert.Equal(t, v.output, r)
	}
}
