package hashset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewHashSet()

	assert.Equal(t, true, s.Empty(), "TestNewSet should be empty")
}

func TestNewSetWithTwoItem(t *testing.T) {
	s := NewHashSet()

	s.Add(1)
	s.Add(2)
	assert.Equal(t, 2, s.Size(), "TestNewSetWithTwoItem expecting 2 items in set")
}
