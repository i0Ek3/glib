package mutex

import (
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	m := NewMutex()
	m.Lock()
	_ = m.Locked()
	m.Unlock()
	_ = m.Locked()
}

func BenchmarkLock(b *testing.B) {
	m := NewMutex()
	for i := 0; i < b.N; i++ {
		m.Lock()
		_ = i * i
		m.Unlock()
	}
}

func TestTryLock(t *testing.T) {
	m := NewMutex()
	ok := m.TryLock()
	n := 0
	if ok {
		_ = n + n
		m.Unlock()
	}
}

func BenchmarkTryLock(b *testing.B) {
	m := NewMutex()
	for i := 0; i < b.N; i++ {
		ok := m.TryLock()
		if ok {
			_ = i * i
			m.Unlock()
		}
	}
}

func TestLockTimeout(t *testing.T) {
	m := NewMutex()
	m.Lock()
	_ = m.LockTimeout(time.Second)
	m.Unlock()
	_ = m.UnlockTimeout(time.Second)
}

func BenchmarkLockWithTimeout(b *testing.B) {
	m := NewMutex()
	for i := 0; i < b.N; i++ {
		m.Lock()
		_ = m.LockTimeout(time.Second)
		_ = i * i
		m.Unlock()
		_ = m.UnlockTimeout(time.Second)
	}
}
