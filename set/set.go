package set

// Set ...
type Set struct {
	items map[interface{}]struct{}
}

// Add ...
func (s *Set) Add(element interface{}) bool {
	if !s.Has(element) {
		if s.items == nil {
			s.items = make(map[interface{}]struct{})
		}
		s.items[element] = struct{}{}
		return true
	}
	return false
}

// Remove ...
func (s *Set) Remove(element interface{}) bool {
	if s.Has(element) {
		delete(s.items, element)
		return true
	}
	return false
}

// Has ...
func (s *Set) Has(element interface{}) bool {
	_, has := s.items[element]
	return has
}

// Clear ...
func (s *Set) Clear() {
	s.items = make(map[interface{}]struct{})
}

// Size ...
func (s *Set) Size() int {
	if s.items == nil {
		return 0
	}
	cnt := 0
	for range s.items {
		cnt++
	}
	return cnt
}

// Values ...
func (s *Set) Values() []interface{} {
	var v []interface{}
	for k := range s.items {
		v = append(v, k)
	}
	return v
}
