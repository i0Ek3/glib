package math

import(
    "math"
)

type Int interface {
    ~int | ~uint
}

func max[T Int](vals ...T) T {
    maxv := T(math.Inf(-1))
    for _, v := range vals {
        if v > maxv {
            maxv = v
        }
    }
    return maxv
}

func MaxFromNumbers[T Int](vals ...T) T {
    return max(vals...)
}

func Max[T Int](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func Abs[T Int](a T) T {
    if a > 0 {
        return a
    }
    return -a
}
