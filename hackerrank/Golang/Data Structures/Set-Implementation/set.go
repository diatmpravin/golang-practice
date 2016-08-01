package set

type Set struct {
	set map[interface{}]bool
}

func NewSet() *Set {
	return &Set{make(map[interface{}]bool)}
}

func (set *Set) Size() int {
	return len(set.set)
}

func (set *Set) Add(i interface{}) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found //False if it existed already
}
