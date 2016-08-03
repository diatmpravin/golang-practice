package hashset

type Set struct {
	set map[interface{}]bool
}

func NewHashSet() *Set {
	return &Set{make(map[interface{}]bool)}
}
