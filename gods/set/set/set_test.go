package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet()

	assert.Equal(t, true, s.Empty(), "New set should be empty")
}

func TestAddItemToEmptySet(t *testing.T) {
	s := NewSet()

	s.Add(1)
	assert.Equal(t, 1, s.Size(), "Expecting set length to be 1")
}

func TestAddMultipleItemsToSet(t *testing.T) {
	s := NewSet()

	s.Add(1, 2)
	assert.Equal(t, 2, s.Size(), "Expecting set length to be 2")

}

func TestNoDuplicateItemsToSet(t *testing.T) {
	s := NewSet()

	s.Add(1, 2, 2)
	assert.Equal(t, 2, s.Size(), "Expecting set length to be 1")

}

func TestMultipleDataTypeToSet(t *testing.T) {
	s := NewSet()

	s.Add(1)
	s.Add("pravin")
	assert.Equal(t, 2, s.Size(), "Expecting set length to be 2")

}

func TestItemContainToSet(t *testing.T) {
	s := NewSet()

	s.Add(1)
	s.Add("pravin")

	assert.Equal(t, true, s.Contain("pravin"), "Expecting set should contain pravin")
	assert.Equal(t, true, s.Contain(1), "Expecting set should contain 1")

}

func TestItemNotContainToSet(t *testing.T) {
	s := NewSet()

	s.Add(1)
	s.Add("pravin")

	assert.Equal(t, false, s.Contain("ram"), "Expecting set should contain pravin")
	assert.Equal(t, false, s.Contain(3), "Expecting set should contain 1")

}

func TestRemoveItemFromSet(t *testing.T) {
	s := NewSet()

	s.Add(1)
	s.Add(2)
	s.Add("pravin")
	s.Add("deo")

	s.Remove(2)

	assert.Equal(t, 3, s.Size(), "Expecting set length to be 3, as one item is removed")
	assert.Equal(t, false, s.Contain(2), "Set should not contain 2 any more")

	s.Remove("pravin")

	assert.Equal(t, 2, s.Size(), "Expecting set length to be 2, as two item is removed")
	assert.Equal(t, false, s.Contain("pravin"), "Set should not contain 'pravin' any more")

}

func TestClearSet(t *testing.T) {
	s := NewSet()

	s.Add(1)
	s.Add(2)
	s.Add("pravin")
	s.Add("deo")

	assert.Equal(t, 4, s.Size(), "Expecting set length to be 4")

	s.Clear()
	assert.Equal(t, true, s.Empty(), "Expecting set should be empty, as set got cleared")

}

func TestUnionOfTwoSets(t *testing.T) {
	s1 := NewSet()
	s1.Add(1, 2, 3)

	s2 := NewSet()
	s2.Add(1, 3, 4, 5)

	s3 := s1.Union(s2)

	assert.Equal(t, 5, s3.Size(), "Expecting set length to be 5")
	assert.Equal(t, true, s3.Contain(3), "Set should contain 3")
	assert.Equal(t, true, s3.Contain(4), "Set should contain 4")
	assert.Equal(t, true, s3.Contain(5), "Set should contain 5")

}

func TestIntersectionOfTwoSets(t *testing.T) {
	s1 := NewSet()
	s1.Add(1, 2, 3)

	s2 := NewSet()
	s2.Add(1, 3, 4, 5)

	s3 := s1.Intersection(s2)
	assert.Equal(t, 2, s3.Size(), "Expecting set length to be 2")
	assert.Equal(t, true, s3.Contain(1), "Set should contain 1")
	assert.Equal(t, true, s3.Contain(3), "Set should contain 3")
}
