package transfer

import (
	"reflect"
	"unsafe"
)

// s2b converts string to []byte
func str2byte(s string) []byte {
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	byt := reflect.SliceHeader{
		Data: str.Data,
		Len:  str.Len,
		Cap:  str.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&byt))
}

// b2s converts []byte to string
func byte2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
