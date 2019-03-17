package converter

import (
	"reflect"
	"unsafe"
)

// String2ByteSlice fast and no memory alloction converter from string to []byte
func String2ByteSlice(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// ByteSlice2String fast and no memory allocation converter from []byte to string
func ByteSlice2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
