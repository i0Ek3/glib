package sema

type Locker interface {
	Lock()
	Unlock()
	TryLock() bool
}

// Semaphore indicates a semaphore
type Semaphore struct {
	Locker
	ch chan struct{}
}

// NewSemaphore returns a Locker interface
func NewSemaphore(capa int) Locker {
	if capa <= 0 {
		capa = 1
	}
	return &Semaphore{
		// create a buffered channel which capacity is capa
		ch: make(chan struct{}, capa),
	}
}

// Lock acquires a resource
func (s *Semaphore) Lock() {
	s.ch <- struct{}{}
}

// Unlock releases a resource
func (s *Semaphore) Unlock() {
	<-s.ch
}

// TryLock tries to acquire a resource
func (s *Semaphore) TryLock() bool {
	select {
	case s.ch <- struct{}{}:
		return true
	default:
	}
	return false
}
