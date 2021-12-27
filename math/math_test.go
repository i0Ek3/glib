package math

import(
    "testing"
)

func TestMaxFromNumbers(t *testing.T) {
    got := MaxFromNumbers(1, 2, 3, 5, 3)
    want := 5

    if got != want {
        t.Errorf("got %q not equal want %q", got, want)
    }
}

func TestMax(t *testing.T) {
    got := Max(1, 2)
    want := 2

    if got != want {
        t.Errorf("got %q not equal want %q", got, want)
    }
}

func TestAbs(t *testing.T) {
    got := Abs(-4)
    want := 4

    if got != want {
        t.Errorf("got %q not equal want %q", got, want)
    }
}
