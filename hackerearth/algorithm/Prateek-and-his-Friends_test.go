package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFriendList(t *testing.T) {
	var friendListTest = []struct {
		n, x   int
		cost   []int
		output string
	}{
		{0, 12, []int{}, "NO"},
		{5, 0, []int{1, 3, 4, 5, 2}, "NO"},
		{5, 1, []int{1, 3, 4, 5, 2}, "YES"},
		{5, 12, []int{1, 3, 4, 5, 2}, "YES"},
		{5, 12, []int{1, 3, 4, 9, 2}, "NO"},
		{5, 12, []int{1, 3, 4, 10, 2}, "YES"},
	}

	for _, v := range friendListTest {
		result := friendList(v.n, v.x, v.cost)
		assert.Equal(t, v.output, result)
	}
}
