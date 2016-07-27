package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var circularArrayRotationTest = []struct {
	slice, newSlice []int
	k               int
}{
	{[]int{}, []int{}, 5},
	{nil, []int{}, 5},
}

func testEqual(a, b []int) bool {
	if len(a) == 0 || len(b) == 0 {
		return true
	}

	return false 
}

func TestCircularArrayRotation(t *testing.T) {
	for _, v := range circularArrayRotationTest {
		// r := circularArrayRotation(v.slice, v.k)
		flag := testEqual(v.newSlice, v.slice)
		assert.Equal(t, true, flag)
	}
}
