package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqualStack(t *testing.T) {
	var equalStackTest = []struct {
		slice1, slice2, slice3 []int
		result                 int
	}{
		{[]int{3, 2, 1, 1, 1}, []int{4, 3, 2}, []int{1, 1, 4, 1}, 5},
	}

	for _, v := range equalStackTest {
		r := equalStack(v.slice1, v.slice2, v.slice3)
		assert.Equal(t, v.result, r)
	}
}
