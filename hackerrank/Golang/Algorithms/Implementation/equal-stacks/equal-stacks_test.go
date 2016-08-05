package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqualStack(t *testing.T) {
	var equalStackTest = []struct {
		slice1, slice2, slice3 []int64
		result                 int64
	}{
		{[]int64{3, 2, 1, 1, 1}, []int64{4, 3, 2}, []int64{1, 1, 4, 1}, 5},
		{[]int64{1, 1, 1, 1, 1}, []int64{3, 2}, []int64{1, 3, 1}, 5},
	}

	for _, v := range equalStackTest {
		r := equalStack(v.slice1, v.slice2, v.slice3)
		assert.Equal(t, v.result, r)
	}
}
