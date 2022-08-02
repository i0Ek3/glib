package lock

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/petermattis/goid"
)

// RecurMutex use goroutine id to implement ReentrantLock
type RecurMutex struct {
	sync.Mutex
	owner int64
	times int32
}

var id = goid.Get()

func (m *RecurMutex) LockByID(id int64) {
	m.lockByID(id)
}

func (m *RecurMutex) lockByID(id int64) {
	if atomic.LoadInt64(&m.owner) == id {
		m.times++
		return
	}
	m.Lock()
	atomic.StoreInt64(&m.owner, id)
	m.times = 1
}

func (m *RecurMutex) UnlockByID(id int64) {
	m.unlockByID(id)
}

func (m *RecurMutex) unlockByID(id int64) {
	if atomic.LoadInt64(&m.owner) != id {
		panic(fmt.Sprintf("Unexpected owner(%d): %d!", m.owner, id))
	}
	m.times--
	if m.times != 0 {
		return
	}
	atomic.StoreInt64(&m.owner, -1)
	m.Unlock()
}

// TokenRecurMutex use token to implement ReentrantLock
type TokenRecurMutex struct {
	sync.Mutex
	token int64
	times int32
}

func (m *TokenRecurMutex) LockByToken(token int64) {
	m.lockByToken(token)
}

func (m *TokenRecurMutex) lockByToken(token int64) {
	if atomic.LoadInt64(&m.token) == token {
		m.times++
		return
	}
	m.Lock()
	atomic.StoreInt64(&m.token, token)
	m.times = 1
}

func (m *TokenRecurMutex) UnlockByToken(token int64) {
	m.unlockByToken(token)
}

func (m *TokenRecurMutex) unlockByToken(token int64) {
	if atomic.LoadInt64(&m.token) != token {
		panic(fmt.Sprintf("Unexpected token(%d): %d!", m.token, token))
	}
	m.times--
	if m.times != 0 {
		return
	}
	atomic.StoreInt64(&m.token, 0)
	m.Unlock()
}
