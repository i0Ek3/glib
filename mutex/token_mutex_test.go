package mutex

import (
	"fmt"
	"testing"
)

const token = 239472631907

func TestLockByToken(t *testing.T) {
	m := new(TokenRecurMutex)
	m.LockByToken(token)
	fmt.Println("Test LockByToken")
	m.LockByToken(token)
	m.UnlockByToken(token)
}

func BenchmarkLockByToken(b *testing.B) {
	m := new(TokenRecurMutex)
	for i := 0; i < b.N; i++ {
		m.LockByToken(token)
		m.LockByToken(token)
		m.UnlockByToken(token)
	}
}
