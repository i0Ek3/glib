package lock

import (
	"fmt"
	"testing"
)

func TestLockByID(t *testing.T) {
	m := new(RecurMutex)
	m.LockByID(id)
	fmt.Println("doing sth...")
	m.LockByID(id)
	m.UnlockByID(id)
}

const token = 239472631907

func TestLockByToken(t *testing.T) {
	m := new(TokenRecurMutex)
	m.LockByToken(token)
	fmt.Println("doing sth...")
	m.LockByToken(token)
	m.UnlockByToken(token)
}
