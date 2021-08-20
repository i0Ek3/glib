package transfer

func bool2int(b bool) int {
    if b {
        return 1
    }
    return 0
}

func int2bool(i int) bool {
    return i != 0
}
