package minit

//import "fmt"

// TODO: try to use Go's new feature generic to improve it, make(map[kT]vT)
func minit(m map[int]int) (res map[int]int) {
    if m == nil {
        res = make(map[int]int)
    }
    return res
}

func Minit(m map[int]int) (res map[int]int) {
    res = minit(m)
    return
}
