package transfer

func B2I(b bool) int {
    if b {
        return 1
    }
    return 0
}

func I2B(i int) bool {
    return i != 0
}
