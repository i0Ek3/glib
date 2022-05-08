package lock

import (
    "fmt"
    "testing"
)

const token = 239472631907

func TestLockByToken(t *testing.T) {
    m := new(TokenRecurMutex)
    m.LockByToken(token)
    fmt.Println("doing sth...")
    m.LockByToken(token)
    m.UnlockByToken(token)
}
