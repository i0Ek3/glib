package transfer

import (
	"github.com/i0Ek3/asrt"
	"testing"
)

func TestS2B(t *testing.T) {
    oristr := "orignal string"
    got := S2B(oristr)
    want := []byte(oristr)

    asrt.Equal(got, want)
}

func TestB2S(t *testing.T) {
    for i := 'A'; i <= 'Z'; i++ {
        oribyt := []byte{byte(i)}
        got := B2S(oribyt)
        want := string(oribyt)
        asrt.Equal(got, want)
    } 
}

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

func TestLower(t *testing.T) {
	for i := byte('A'); i <= byte('Z'); i++ {
		got := Lower(i)
		if got != i+32 {
			t.Errorf("got %v != want %v", got, i+32)
		}
	}
}
