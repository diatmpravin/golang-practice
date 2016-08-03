package set

type Set struct {
	set map[interface{}]bool
}

func NewSet() *Set {
	return &Set{make(map[interface{}]bool)}
}

func (s Set) Empty() bool {
	return len(s.set) == 0
}

func (s Set) Size() int {
	return len(s.set)
}

func (s *Set) Contain(i interface{}) bool {
	_, ok := s.set[i]
	return ok
}

func (s *Set) Remove(i interface{}) {
	delete(s.set, i)
}

func (s *Set) Clear() {
	s.set = make(map[interface{}]bool)
}

func (s1 *Set) Union(s2 *Set) *Set {
	for k, _ := range s2.set {
		s1.Add(k)
	}

	return s1
}

func (s1 *Set) Intersection(s2 *Set) *Set {
	newSet := NewSet()
	for k, _ := range s2.set {

		if s1.Contain(k) {
			newSet.Add(k)
		}
	}

	return newSet
}

func (s *Set) Add(i ...interface{}) {
	for _, v := range i {
		_, ok := s.set[v]

		if !ok {
			s.set[v] = true
		}
	}

}
