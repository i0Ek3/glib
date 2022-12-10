package mutex

import (
	"fmt"
	"testing"
)

func TestLockByID(t *testing.T) {
	m := new(RecurMutex)
	m.LockByID(id)
	fmt.Println("Test LockByID")
	m.LockByID(id)
	m.UnlockByID(id)
}

func BenchmarkLockByID(b *testing.B) {
	m := new(RecurMutex)
	for i := 0; i < b.N; i++ {
		m.lockByID(id)
		m.lockByID(id)
		m.unlockByID(id)
	}
}
