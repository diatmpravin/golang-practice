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

	assert.Equal(t, true, aSet.Contains("pravin"), "Set should contains 'pravin' element")
	assert.Equal(t, true, aSet.Contains("ankit"), "Set should contains 'ankit' element")
	assert.Equal(t, true, aSet.Contains(1), "Set should contains '1' element")
}

func TestRemoveSet(t *testing.T) {
	aSet := NewSet()

	aSet.Add(1)
	aSet.Add("pravin")
	aSet.Add("ankit")

	aSet.Remove("pravin")

	assert.Equal(t, 2, aSet.Size(), "Set size should be 2 as, one elment 'pravin' is removed")
	assert.Equal(t, true, aSet.Contains(1), "Set should only conatins 1 and 'ankit'")
	assert.Equal(t, true, aSet.Contains("ankit"), "Set should only conatins 1 and 'ankit'")

	aSet.Remove(1)
	aSet.Remove("ankit")
	assert.Equal(t, 0, aSet.Size(), "TestRemoveSet should have size ZERO as remaing 1 and 'ankit' is also removed")
}

func TestConainSet(t *testing.T) {
	aSet := NewSet()

	aSet.Add(1)
	aSet.Add("pravin")
	aSet.Add("ankit")

	assert.Equal(t, true, aSet.Contains("pravin"), "TestContainSet should contain 'pravin'")
	assert.Equal(t, true, aSet.Contains(1), "TestContainSet should contain 1")
	assert.Equal(t, true, aSet.Contains("ankit"), "TestContainSet should contain 'ankit'")

	aSet.Remove("pravin")
	assert.Equal(t, false, aSet.Contains("pravin"), "TestContainSet should contain 'pravin'")

	aSet.Add("avi")
	assert.Equal(t, true, aSet.Contains("avi"), "TestContainSet should contain 'avi'")

}
