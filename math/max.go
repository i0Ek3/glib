package math

import(
    "math"
)

func max(vals... int) int {
    maxv := int(math.Inf(-1))
    for _, v := range vals {
        if v > maxv {
            maxv = v
        }
    }
    return maxv
}

func Max(vals... int) int {
    return max(vals...)
}
