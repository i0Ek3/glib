package design

import (
	"testing"
)

var in Instance

func TestIntInstance(t *testing.T) {
	want := 1
	got := in.ReturnIntInstance()
	if want != got {
		t.Errorf("wrong!")
	}
}

func TestPointerInstance(t *testing.T) {
	want := (*int)(nil)
	got := in.ReturnPointerInstance()
	if want != got {
		t.Errorf("wrong!")
	}
}

func BenchmarkIntInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go in.ReturnIntInstance()
	}
}

func BenchmarkPointerInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go in.ReturnPointerInstance()
	}
}
