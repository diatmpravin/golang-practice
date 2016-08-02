package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSet(t *testing.T) {
	aSet := NewSet()

	assert.Equal(t, 0, aSet.Size(), "A new set size should be ZERO")
}

func TestAddSet(t *testing.T) {
	aSet := NewSet()

	aSet.Add(1)
	aSet.Add("pravin")
	aSet.Add("ankit")

	assert.Equal(t, 3, aSet.Size(), "After adding 3 item, set size should be THREE")
}

func TestNoDuplicateItem(t *testing.T) {
	aSet := NewSet()

	aSet.Add(1)
	aSet.Add("pravin")
	aSet.Add("ankit")
	aSet.Add(1)

	assert.Equal(t, 3, aSet.Size(), "Set size should be 3 as 1 elements is duplicate")
}
