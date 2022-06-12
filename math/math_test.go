package math

import(
    "testing"
)

func TestTmax(t *testing.T) {
    got := Tmax(1, 2, 3, 5, 3)
    want := 5

    if got != want {
        t.Errorf("got %q not equal want %q", got, want)
    }
}

func TestTmin(t *testing.T) {
    got := Tmin(1, 2, 4, 6)
    want := 1

    if got != want {
        t.Errorf("got %q not equal want %q", got, want)
    }
}

func TestTabs(t *testing.T) {
    got := Tabs(-4, 0, 5, -1)
    want := []int{4, 0, 5, 1}

    sum := 0
    for _, v1 := range got {
        sum += v1
    }

    for _, v2 := range want {
        sum -= v2
    }

    if sum != 0 {
        t.Errorf("got %q not equal want %q", got, want)
    }
}
