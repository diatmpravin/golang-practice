package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var equalStackTest = []struct {
	arr1, arr2, arr3 []int
	result           int
}{
	{[]int{}, []int{}, []int{}, 0},
	{[]int{1}, []int{1}, []int{1}, 1},
	{[]int{1, 1, 1, 1, 1}, []int{1, 2, 2}, []int{4, 1}, 5},
	{[]int{3, 2, 1, 1, 1}, []int{4, 3, 2}, []int{1, 1, 4, 1}, 5},
}

func TestEqualStack(t *testing.T) {
	for _, v := range equalStackTest {
		r := equalStack(v.arr1, v.arr2, v.arr3)
		assert.Equal(t, v.result, r)
	}
}
