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

var arraySumWrongTest = []struct {
	arr []int
	sum int
}{
	{[]int{}, 3},
	{[]int{0}, 2},
	{[]int{0, 0}, 1},
	{[]int{0, 0, 0}, 4},
	{[]int{0, 1, 0}, 2},
	{[]int{1, 2, 3}, 8},
}

func TestArraySum(t *testing.T) {
	for _, v := range arraySumTest {
		r := arraySum(v.arr)
		if assert.NotNil(t, r) {
			assert.Equal(t, v.sum, r, "The sum should be: %d", v.sum)
		}
	}
}

func TestArraySumNotEqual(t *testing.T) {
	for _, v := range arraySumWrongTest {
		r := arraySum(v.arr)
		assert.NotEqual(t, v.sum, r, "The sum should not be: %d", v.sum)
	}
}
