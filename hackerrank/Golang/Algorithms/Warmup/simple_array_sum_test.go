package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var arraySumTest = []struct {
	arr []int
	sum int
}{
	{[]int{}, 0},
	{[]int{0}, 0},
	{[]int{0, 0}, 0},
	{[]int{0, 0, 0}, 0},
	{[]int{0, 1, 0}, 1},
	{[]int{1, 2, 3}, 6},
}

func TestArraySum(t *testing.T) {
	for _, v := range arraySumTest {
		r := arraySum(v.arr)
		assert.Equal(t, v.sum, r)
	}
}
