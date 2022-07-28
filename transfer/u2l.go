package transfer

func Lower(c byte) byte {
    return c | ('x' - 'X')
}
