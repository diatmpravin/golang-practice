package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var circularArrayRotationTest = []struct {
	slice, newSlice []int
	k               int
	eq              bool
	index           int
	r               int
}{
	{[]int{}, []int{}, 5, true, 8, 0},
	{[]int{0}, []int{}, 5, false, 2, 0},
	{nil, []int{1}, 5, false, 0, 0},
	{nil, nil, 5, true, 0, 0},
	{[]int{1, 2, 3}, []int{2, 3, 1}, 2, true, 1, 3},
	{[]int{1, 2, 3}, []int{1, 2, 3}, 3, true, 2, 3},
}

func testEqual(a, b []int) bool {
	fmt.Println(a, b)
	if len(a) == 0 && len(b) == 0 {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestCircularArrayRotation(t *testing.T) {
	for _, v := range circularArrayRotationTest {
		r := circularArrayRotation(v.slice, v.k)
		flag := testEqual(v.newSlice, r)
		assert.Equal(t, v.eq, flag)
	}
}

func TestFindElement(t *testing.T) {
	for _, v := range circularArrayRotationTest {
		if len(v.newSlice) > 0 && v.eq {
			r := findElement(v.newSlice, v.index)
			assert.Equal(t, v.r, r)
		}
	}
}