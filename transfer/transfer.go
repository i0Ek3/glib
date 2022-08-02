package transfer

import (
	"reflect"
	"unsafe"
)

func B2I(b bool) int {
	if b {
		return 1
	}
	return 0
}

func I2B(i int) bool {
	return i != 0
}

func Lower(c byte) byte {
	return c | ('x' - 'X')
}

func str2byte(s string) []byte {
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	byt := reflect.SliceHeader{
		Data: str.Data,
		Len:  str.Len,
		Cap:  str.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&byt))
}

func S2B(s string) []byte {
	return str2byte(s)
}

func byte2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func B2S(b []byte) string {
	return byte2str(b)
}

func ReverseWithInterval(b []byte, left, right int) string {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
	return string(b)
}

func Reverse(res []int) []int {
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
