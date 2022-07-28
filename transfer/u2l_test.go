package transfer

import (
    "testing"
)

func TestLower(t *testing.T) {
    for i := byte('A'); i <= byte('Z'); i++ {
        got := Lower(i)
        if got != i+32 {
            t.Errorf("got %v != want %v", got, i+32)
        }
    }
}
