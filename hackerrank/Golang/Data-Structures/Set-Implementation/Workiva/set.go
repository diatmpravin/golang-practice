package set

import (
	"fmt"
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

func NewSet(items ...interface{}) *Set {
	set := pool.Get().(*Set)
	fmt.Printf("When New set Created:------> %+v \n", set)
	for _, item := range items {
		set.items[item] = struct{}{}
	}

	fmt.Printf("At the time of return:-----> %+v \n", set)
	return set
}
