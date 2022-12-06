package mutex

import (
	"time"
)

// Mutex indicates the mutex in Go which implement by channel
type Mutex struct {
	ch chan struct{}
}

// NewMutex create a mutex object
func NewMutex() *Mutex {
	mu := &Mutex{make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

// Lock requires a lock by receive a value from channel
func (m *Mutex) Lock() {
	<-m.ch
}

// Unlock releases a lock by send a value to channel
func (m *Mutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("Cannot unlock unlocked mutex!")
	}
}

// TryLock tries to require a lock and return true else false
func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}

	return false
}

// LockTimeout sets a timeout for lock
func (m *Mutex) LockTimeout(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-m.ch:
		timer.Stop()
		return true
	case <-timer.C:
	}

	return false
}

// UnlockTimeout sets a timeout for unlock
func (m *Mutex) UnlockTimeout(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case m.ch <- struct{}{}:
		timer.Stop()
		return true
	case <-timer.C:
	}

	return false
}

// Locked returns whether a lock was held
func (m *Mutex) Locked() bool {
	return len(m.ch) == 0
}
