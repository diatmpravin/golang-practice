package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet()

	assert.Equal(t, 0, len(s.items), "TestNewSet expecting empty set")
}

func TestNewSetWithOneItem(t *testing.T) {
	s := NewSet(1)

	assert.Equal(t, 1, len(s.items), "TestNewSetWithOneItem should have 1 item in set")
}

func TestNewSetWithThreeItem(t *testing.T) {
	s := NewSet(1, 2, 3, 3)

	assert.Equal(t, 3, len(s.items), "TestNewSetWithThreeItem should have 3 item in set")
}

func TestAddItems(t *testing.T) {
	s := NewSet()

	s.Add(`test`)
	s.Add(`test1`)
	assert.Equal(t, 2, len(s.items), "TestAddItems should have 2 items in set")
}
