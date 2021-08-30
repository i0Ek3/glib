<h1 align="center">glib👋</h1>
<p>
</p>

> glib: A lots of Go utils for myself.

## Content

- minit
- math
- set
- transfer


## Getting Started

### Installation

`go get github.com/i0Ek3/glib`


### Usage

```Go
import (
    mi "github.com/i0Ek3/glib/minit"
    tr "github.com/i0Ek3/glib/transfer"
    mx "github.com/i0Ek3/glib/math"
    set "github.com/i0Ek3/glib/set"
    ...
)

func main() {
    // for minit
    m := make(map[int]int) // TODO: Apply Go generic
    mi.Minit(m)

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
}

```



## Contributing

PRs and Issues are also welcome.


## Show your support

Give a ⭐️ if this project helped you!

