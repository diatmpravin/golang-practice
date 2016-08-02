package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSet(t *testing.T) {
	aSet := NewSet()

	assert.Equal(t, 0, aSet.Size(), "A new set size should be ZERO")
}
