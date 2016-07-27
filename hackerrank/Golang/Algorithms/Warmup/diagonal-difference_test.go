package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var diagonalDifferenceTest = []struct {
	s    [][]int
	diff int
}{
	{[][]int{{}, {}}, 0},
	{[][]int{{0, 0}, {0, 0}}, 0},
	{[][]int{{1, 2}, {3, 4}}, 0},
}

func TestDiagonalDifference(t *testing.T) {
	for _, v := range diagonalDifferenceTest {
		r := diagonalDifference(len(v.s), v.s)
		if assert.NotNil(t, r) {
			assert.Equal(t, v.diff, r)
		}
	}
}
