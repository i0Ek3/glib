
# glib

`glib` is a series of small libraries written in Golang, designed to help developers focus their development work on more noteworthy areas. You can use the libraries in glib to help you develop faster in your projects, welcome to try.


## Features

- Support ReentrantLock
- Support thread-safe Set
- Support Mutex/Once version Singleton
- Support type transfer
- Support math operation

## Install

```Go
go get -u github.com/i0Ek3/glib
```


## Usage/Examples

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

Contributions are always welcome!

## Support

For support, email i0Ek3@protonmail.ch.


## License

[MIT](https://choosealicense.com/licenses/mit/)


## Credit

[<img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png" style="zoom:25%;" />](https://jb.gg/OpenSourceSupport)

