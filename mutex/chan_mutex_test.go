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
	_ = m.TryLock()
}

func BenchmarkTryLock(b *testing.B) {
	m := NewMutex()
	for i := 0; i < b.N; i++ {
		_ = m.TryLock()
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
