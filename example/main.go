package main

import (
	"fmt"

	"github.com/i0Ek3/glib/lock"
	"github.com/i0Ek3/glib/math"
	"github.com/i0Ek3/glib/pool"
	"github.com/i0Ek3/glib/set"
	"github.com/i0Ek3/glib/transfer"
	"github.com/petermattis/goid"
)

func main() {
	o := fmt.Println
	// for transfer
	o(transfer.I2B(100))
	o(transfer.B2I(true))
	o(transfer.S2B("glib"))

	// for math
	o(math.Tmax([]int{1, 2, 3, 4, 5})) // 5
	o(math.Tmin([]int{1, 2, 3, 4}))    // 1
	o(math.Tabs([]int{-1, -2, -3}))    // 1, 2, 3
	o(math.Ttrop(true, "glib", "i0Ek3"))

	// for set
	set.NewSet(1, 2, 3)
	s := set.Set{}
	s.Add(1)
	s.Remove(3)
	s.Clear()
	s.Query(2)
	s.GetSize()

	// for pool
	pool := pool.New(4)
	pool.NewTask(func() {
		fmt.Println("this is pool package")
	})

	// for lock
	// lock by goroutine id
	rL := lock.RecurMutex{}
	id := goid.Get()
	rL.LockByID(id)
	rL.LockByID(id)
	rL.UnlockByID(id)

	// lock by token
	tL := lock.TokenRecurMutex{}
	tL.LockByToken(123456789)
	tL.LockByToken(123456789)
	tL.UnlockByToken(123456789)
}
