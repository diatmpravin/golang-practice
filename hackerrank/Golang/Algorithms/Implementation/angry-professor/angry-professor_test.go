package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClassCancelled(t *testing.T) {
	var classCancelledTest = []struct {
		s      []int
		k      int
		result string
	}{
		{[]int{}, 0, "NO"},
		{[]int{}, 0, "NO"},
		{[]int{-1, -3, 4, 2}, 3, "YES"},
		{[]int{0, -1, 2, 1}, 2, "NO"},
	}

	for _, v := range classCancelledTest {
		r := classCancelled(v.s, v.k)
		assert.Equal(t, v.result, r)
	}
}
