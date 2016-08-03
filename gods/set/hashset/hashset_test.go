package hashset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewHashSet()

	assert.Equal(t, 0, len(s.set), "TestNewSet should be empty")
}
