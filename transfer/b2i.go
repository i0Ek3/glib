package transfer

func bool2int(b bool) int {
    if b {
        return 1
    }
    return 0
}

func B2I(b bool) int {
    return bool2int(b)
}

func int2bool(i int) bool {
    return i != 0
}

func I2B(i int) bool {
    return int2bool(i)
}
