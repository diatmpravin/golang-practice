package hashset

type Set struct {
	set map[interface{}]bool
}

func NewHashSet() *Set {
	return &Set{make(map[interface{}]bool)}
}

func (s *Set) Add(i ...interface{}) {
	for _, v := range i {
		_, found := s.set[v]
		if !found {
			s.set[v] = true
		}
	}
}

func (s Set) Size() int {
	return len(s.set)
}

func (s Set) Empty() bool {
	return len(s.set) == 0
}
