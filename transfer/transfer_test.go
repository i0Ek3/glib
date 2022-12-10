package transfer

import (
	"testing"

	"github.com/i0Ek3/asrt"
)

func TestS2B(t *testing.T) {
	oristr := "orignal string"
	got := S2B(oristr)
	want := []byte(oristr)

	asrt.Equal(got, want)
}

func BenchmarkS2B(b *testing.B) {
	oristr := "orignal string"
	for i := 0; i < b.N; i++ {
		_ = S2B(oristr)
	}
}

func TestB2S(t *testing.T) {
	for i := 'A'; i <= 'Z'; i++ {
		oribyt := []byte{byte(i)}
		got := B2S(oribyt)
		want := string(oribyt)
		asrt.Equal(got, want)
	}
}

func BenchmarkB2S(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 'A'; j <= 'Z'; j++ {
			oribyt := []byte{byte(j)}
			_ = B2S(oribyt)
		}
	}
}

func TestB2R(t *testing.T) {
	for i := byte('A'); i <= ('Z'); i++ {
		_ = B2R(i)
	}
}

func BenchmarkB2R(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := byte('A'); j <= ('Z'); j++ {
			_ = B2R(j)
		}
	}
}

func TestR2B(t *testing.T) {
	for i := 'A'; i <= 'Z'; i++ {
		_ = R2B(i)
	}
}

func BenchmarkR2B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 'A'; j <= 'Z'; j++ {
			_ = R2B(j)
		}
	}
}

func TestReverseWithInterval(t *testing.T) {
	b := "hello"
	got := ReverseWithInterval([]byte(b), 0, 4)
	want := string(b)

	asrt.Equal(got, want)
}

func BenchmarkReverseWithInterval(b *testing.B) {
	s := "hello"
	for i := 0; i < b.N; i++ {
		_ = ReverseWithInterval([]byte(s), 0, 4)
	}
}

func TestReverse(t *testing.T) {
	res := []int{1, 2, 3, 4, 5}
	got := Reverse(res)
	want := []int{5, 4, 3, 2, 1}

	asrt.Equal(got, want)
}

func BenchmarkReverse(b *testing.B) {
	res := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		_ = Reverse(res)
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

func BenchmarkLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := byte('A'); j <= byte('Z'); j++ {
			_ = Lower(j)
		}
	}
}
