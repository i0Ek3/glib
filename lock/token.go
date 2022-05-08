package lock

import (
	"fmt"
	"sync"
	"sync/atomic"
)

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
