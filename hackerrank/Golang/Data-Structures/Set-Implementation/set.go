package set

type Set struct {
	set map[interface{}]bool
}

func NewSet() *Set {
	return &Set{make(map[interface{}]bool)}
}

func (set Set) Size() int {
	return len(set.set)
}
