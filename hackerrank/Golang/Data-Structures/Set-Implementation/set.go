package set

type Set struct {
	set map[interface{}]bool
}

func NewSet() *Set {
	return &Set{make(map[interface{}]bool)}
}

func (s *Set) Add(i interface{}) {
	_, _ = s.set[i]
	s.set[i] = true
}

func (s *Set) Contains(i interface{}) bool {
	_, ok := s.set[i]
	return ok
}

func (s *Set) Remove(i interface{}) {
	delete(s.set, i)
}

func (s Set) Size() int {
	return len(s.set)
}

func (s *Set) Clear() {
	s.set = make(map[interface{}]bool)
}
