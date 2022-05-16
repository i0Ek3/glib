package mutex

import (
    "time"
    "testing"
)

func TestLock(t *testing.T) {
    m := NewMutex()
    m.Lock()
    ok := m.Locked()
    if ok {
        t.Errorf("Locked!")
    }
    m.Unlock()
    ok = m.Locked()
    if ok {
        t.Errorf("Unlocked!")
    }
}

func TestTryLock(t *testing.T) {
    m := NewMutex()
    ok := m.TryLock()
    if ok {
        t.Errorf("TryLock success!")
    }
    ok = m.TryLock()
    if ok {
        t.Errorf("TryLock success!")
    }
}

func TestLockTimeout(t *testing.T) {
    m := NewMutex()
    m.Lock()
    ok := m.LockTimeout(time.Second)
    if ok {
        t.Errorf("Lock timeout!")
    }
}

func TestUnlockTimeout(t *testing.T) {
    m := NewMutex()
    m.Lock()
    m.Unlock()
    ok := m.UnlockTimeout(time.Second)
    if ok {
        t.Errorf("Unlock timeout!")
    }
}
