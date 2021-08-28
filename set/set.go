package set

// Set defines a set struct which made of map
type Set struct {
    set map[interface{}]struct{}
}

// NewSet creates a set and add vals to this set
func NewSet(vals ...interface{}) *Set {
    s := &Set{}
    s.set = make(map[interface{}]struct{})
    s.Add(vals...)
    return s
}

// Add adds given values into set
func (s *Set) Add(vals ...interface{}) {
    for _, v := range vals {
        s.set[v] = struct{}{}
    }
}

// Remove removes the given value from set
func (s *Set) Remove(vals interface{}) {
    delete(s.set, vals)
}

// Clear clears set to nil
func (s *Set) Clear() {
    s.set = nil
}

// Query checks if the given value exist in set
func (s *Set) Query(vals interface{}) bool {
    _, ok := s.set[vals]
    return ok
}

// GetSize gets the size of set
func (s *Set) GetSize() int {
    return len(s.set)
}
