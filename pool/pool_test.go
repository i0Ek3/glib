package pool

import (
    "testing"
    "fmt"
)

func TestPool(t *testing.T) {
    t.Run("test create task by NewTask()", func(t *testing.T) {
        pool := New(8)
        pool.NewTask(func() {
            fmt.Println("this is a test")
        })
    })
}
