package hashset

type Set struct {
	set map[interface{}]bool
}

func NewHashSet() *Set {
	return &Set{make(map[interface{}]bool)}
}

func (s Set) Size() int {
	return len(s.set)
}

func (s Set) Empty() bool {
	return len(s.set) == 0
}

func (s *Set) Add(i ...interface{}) {
	for _, v := range i {
		_, found := s.set[v]
		if !found {
			s.set[v] = true
		}
	}
}

func (s *Set) Contains(i ...interface{}) bool {
	for _, item := range i {
		if _, contains := s.set[item]; !contains {
			return false
		}
	}
	return true
}

func (s *Set) Remove(i ...interface{}) {
	for _, item := range i {
		delete(s.set, item)
	}
}
