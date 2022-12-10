package sema

import "testing"

func TestLock(t *testing.T) {
	sp := NewSemaphore(1)
	for i := 0; i < 100; i++ {
		sp.Lock()
		_ = i * i
		sp.Unlock()
	}
}

func BenchmarkLock(b *testing.B) {
	sp := NewSemaphore(1)
	for i := 0; i < b.N; i++ {
		sp.Lock()
		_ = i * i
		sp.Unlock()
	}
}

func TestTryLock(t *testing.T) {
	sp := NewSemaphore(1)
	for i := 0; i < 100; i++ {
		ok := sp.TryLock()
		if ok {
			_ = i * i
			sp.Unlock()
		}
	}
}

func BenchmarkTryLock(b *testing.B) {
	sp := NewSemaphore(1)
	for i := 0; i < b.N; i++ {
		ok := sp.TryLock()
		if ok {
			sp.Unlock()
		}
	}
}
