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
