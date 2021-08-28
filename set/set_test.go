package set

import(
    "testing"
)

func TestAdd(t *testing.T) {
    s := NewSet(1, 2, 3)
    num := 4
    s.Add(num)
    got := s.Query(num)
    want := true

    if got != want {
        t.Errorf("---Not Equals---")
    }
}

func TestRemove(t *testing.T) {
    s := NewSet(1, 2, 3)
    s.Remove(1)
    got := s.Query(1)
    want := false

    if got != want {
        t.Errorf("---Not Equals---")
    }
}

func TestClear(t *testing.T) {
    s := NewSet(1, 2, 3)
    s.Clear()
    got := s.GetSize()
    want := 0

    if got != want {
        t.Errorf("---Not Equals---")
    }
}

func TestQuery(t *testing.T) {
    s := NewSet(1, 2, 3)
    got := s.Query(1)
    want := true

    if got != want {
        t.Errorf("---Not Equals---")
    }
}

func TestGetSize(t *testing.T) {
    s := NewSet(1, 2, 3)
    got := s.GetSize()
    want := 3

    if got != want {
        t.Errorf("---Not Equals---")
    }
}
