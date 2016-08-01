package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSet(t *testing.T) {
	a := NewSet()

	assert.Equal(t, 0, a.Size(), "NewSet should start out as an empty set")
}

func TestAddSet(t *testing.T) {
	a := NewSet()

	a.Add(3)
	a.Add("pravin")
	a.Add("ankit")

	assert.Equal(t, 3, a.Size(), "AddSet size should be 3 as, three new items were added")
}