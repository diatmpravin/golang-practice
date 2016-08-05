package set

import (
	"sync"
)

var pool = sync.Pool{}

type Set struct {
	items     map[interface{}]struct{}
	lock      sync.RWMutex
	flattened []interface{}
}

func init() {
	pool.New = func() interface{} {
		return &Set{
			items: make(map[interface{}]struct{}, 10),
		}
	}
}

func (s *Set) Add(items ...interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.flattened = nil
	for _, item := range items {
		s.items[item] = struct{}{}
	}
}

func NewSet(items ...interface{}) *Set {
	set := pool.Get().(*Set)
	for _, item := range items {
		set.items[item] = struct{}{}
	}

	return set
}
