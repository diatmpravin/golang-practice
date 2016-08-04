/*
2
5
2 1 5 3 4
5
2 5 1 3 4

3
Too chaotic

*/

package main 

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewYearChaos(t *testing.T) {
	var newYearChaosTest = []struct {
		r   int
		arr []int
	}{
		{-1, []int{}},
		{-1, []int{0}},
		{-1, []int{1}},
		{0, []int{1, 2}},
		{1, []int{2, 1}},
		{0, []int{1, 2, 3}},
		{1, []int{1, 3, 2}},
		{1, []int{2, 1, 3}},
		{2, []int{3, 1, 2}},
		{-1, []int{4, 1, 2, 3}},
		{2, []int{1, 4, 2, 3}},
		{3, []int{2, 1, 5, 3, 4}},
		{-1, []int{1, 5, 2, 3, 4}},
		{3, []int{2, 5, 1, 3, 4}},
	}

	for _, v := range newYearChaosTest {
		result := NewYearChaos(v.arr)
		assert.Equal(t, v.r, result)
	}
}
