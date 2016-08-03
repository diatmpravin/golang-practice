package set

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInt64Set(t *testing.T) {
	s := NewInt64Set()
	fmt.Println("NewInt64Set: ", s)
	assert.Equal(t, 0, len(s), "Expecting new set with 0 elements")
}
