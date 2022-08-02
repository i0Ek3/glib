package transfer

import (
	"github.com/i0Ek3/asrt"
	"testing"
)

func TestReverseWithInterval(t *testing.T) {
	b := "hello"
	got := ReverseWithInterval([]byte(b), 0, 4)
	want := string(b)

	asrt.Equal(got, want)
}

func TestReverse(t *testing.T) {
	res := []int{1, 2, 3, 4, 5}
	got := Reverse(res)
	want := []int{5, 4, 3, 2, 1}

	asrt.Equal(got, want)
}

func TestS2B(t *testing.T) {
	str := "s2b"
	got := S2B(str)
	want := S2B(string(got))
	if string(got) != string(want) {
		t.Error("X")
	}
}

func TestB2S(t *testing.T) {
	str := "s2b"
	byt := []byte(str)
	got := B2S(byt)
	want := string(byt)
	if got != want {
		t.Error("X")
	}
}

func TestLower(t *testing.T) {
	for i := byte('A'); i <= byte('Z'); i++ {
		got := Lower(i)
		if got != i+32 {
			t.Errorf("got %v != want %v", got, i+32)
		}
	}
}
