package set

import (
    "sync"
)

// T/K for generic use
type (
    T interface{}
    K struct{}
)

// Set defines a set struct which made of map
type Set struct {
    sync.RWMutex
    set  map[T]K
}

// NewSet creates a set and add vals to this set
func NewSet(vals ...T) *Set {
    s := &Set{}
    s.set = make(map[T]K)
    s.Add(vals...)
    return s
}

// Add adds given values into set
func (s *Set) Add(vals ...T) {
    for _, v := range vals {
        s.Lock()
        s.set[v] = K{}
        s.Unlock()
    }
}

// Remove removes the given value from set
func (s *Set) Remove(vals T) {
    delete(s.set, vals)
}

// Clear clears set to nil
func (s *Set) Clear() {
    s.set = nil
}

// Query checks if the given value exist in set
func (s *Set) Query(vals T) bool {
    _, ok := s.set[vals]
    return ok
}

// GetSize gets the size of set
func (s *Set) GetSize() int {
    return len(s.set)
}
