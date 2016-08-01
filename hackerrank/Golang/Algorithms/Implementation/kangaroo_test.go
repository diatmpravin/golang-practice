package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWillTheyMeet(t *testing.T) {
	var willTheyMeetTest = []struct {
		x1, v1, x2, v2 int
		output         string
	}{
		{0, 3, 4, 2, "YES"},
		{0, 2, 5, 3, "NO"},
		{4523, 8092, 9419, 8076, "YES"},
	}

	for _, v := range willTheyMeetTest {
		r := willTheyMeet(v.x1, v.v1, v.x2, v.v2)
		assert.Equal(t, v.output, r)
	}
}
