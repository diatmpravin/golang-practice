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

func TestNewSetAdd(t *testing.T) {
	s := NewHashSet()
	s.Add()
	s.Add(1)
	s.Add(2)
	s.Add(2, 3)
	s.Add()

	if actualValue := s.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := s.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}

func TestSetContains(t *testing.T) {
	s := NewHashSet()
	s.Add(3, 1, 2)
	s.Add(2, 3)
	s.Add()

	if actualValue := s.Contains(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := s.Contains(1); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := s.Contains(1, 2, 3); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := s.Contains(1, 2, 3, 4); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestSetRemove(t *testing.T) {
	s := NewHashSet()
	s.Add(3, 1, 2)
	s.Remove()

	if actualValue := s.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	s.Remove(1)
	if actualValue := s.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	s.Remove(3)
	s.Remove(3)
	s.Remove()
	s.Remove(2)
	if actualValue := s.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}
