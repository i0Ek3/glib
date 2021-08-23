<h1 align="center">glib👋</h1>
<p>
</p>

> glib: A lots of Go utils for myself.

## Getting Started

### Installation

`go get github.com/i0Ek3/glib`


### Usage

```Go
import (
    mi "github.com/i0Ek3/glib/minit"
    tr "github.com/i0Ek3/glib/transfer"
    mx "github.com/i0Ek3/glib/math"
    ...
)

func main() {
    m := make(map[int]int) // TODO: Apply Go generic
    mi.Minit(m)

    tr.I2B(/*int value*/1)
    tr.B2I(/*bool value*/true)
    tr.S2B(/*string*/"hello")
    tr.B2S(/*byte*/'h')

    mx.Max(1, 2, 3, 4, 5)
}

```



## Contributing

PRs and Issues are also welcome.


## Show your support

Give a ⭐️ if this project helped you!

