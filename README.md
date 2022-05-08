<h1 align="center">glib👋</h1>
<p>
</p>

> glib: A lots of Go utils for myself.

## Content

- lock
  - ReentrantLock

- pool
- math
- set
- pattern
- transfer

## Getting Started

### Installation

`go get github.com/i0Ek3/glib`

### Usage

Please check example.

```Go
import (
    "github.com/i0Ek3/glib/pool"
    tr "github.com/i0Ek3/glib/transfer"
    mx "github.com/i0Ek3/glib/math"
    set "github.com/i0Ek3/glib/set"
  	l "github.com/i0Ek3/glib/lock"
  
    //...
)

func main() {
    // for transfer
    tr.I2B(/*int value*/1)
    tr.B2I(/*bool value*/true)
    tr.S2B(/*string*/"hello")
    tr.B2S(/*byte*/'h')

    // for math
    mx.MaxFromNumbers(1, 2, 3, 4, 5) // 5
    mx.Max(1, 2) // 2
    mx.Abs(-1) // 1

    // for set
    set.NewSet(1, 2, 3)
    set.Add(1)
    set.Remove(3)
    set.Clear()
    set.Query(2)
    set.GetSize()

    // for pool
    pool := New(4)
    pool.NewTask(task) // var task func()
  
  	// for lock
    // lock by goroutine id
    id := goid.Get()
    l.LockByID(id)
  	l.LockByID(id)
  	l.UnlockByID(id)
  
  	// lock by token
    l.LockByToken(token)
  	l.LockByToken(token)
  	l.UnlockByToken(token)
} 
```

## Contributing

PRs and Issues are also welcome.

## Show your support

Give a ⭐️ if this project helped you!
