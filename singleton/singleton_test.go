package singleton

import (
	"testing"
)

var (
	ptrMInt PtrMutexInstance
	intMInt IntMutexInstance
	ptrOInt PtrOnceInstance
	intOInt IntegerInstance
)

func TestPtrMutexInstance(t *testing.T) {
	want := (*int)(nil)
	got := ptrMInt.ReturnPtrMutexInstance()
	if want != got {
		t.Errorf("wrong!")
	}
}

func TestPtrOnceInstance(t *testing.T) {
	want := (*int)(nil)
	got := ptrOInt.ReturnPtrOnceInstance()
	if want != got {
		t.Errorf("wrong!")
	}
}

func TestIntMutexInstance(t *testing.T) {
	want := 1
	got := intMInt.ReturnIntMutexInstance()
	if want != got {
		t.Errorf("wrong!")
	}
}

func TestIntOnceInstance(t *testing.T) {
	want := 1
	got := intOInt.ReturnIntOnceInstance()
	if want != got {
		t.Errorf("wrong!")
	}
}

func BenchmarkPtrMutexInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go ptrMInt.ReturnPtrMutexInstance()
	}
}

func BenchmarkIntMutexInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go intMInt.ReturnIntMutexInstance()
	}
}

func BenchmarkPtrOnceInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go ptrOInt.ReturnPtrOnceInstance()
	}
}

func BenchmarkIntOnceInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go intOInt.ReturnIntOnceInstance()
	}
}
