package math

import(
    "testing"
)

func TestMax(t *testing.T) {
    got := Max(1, 2, 3, 5, 3)
    want := 5

    if got != want {
        t.Errorf("got %q not equal want %q", got, want)
    }
}
