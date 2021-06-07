package libuseful

import (
	"reflect"
	"unsafe"
)

func DumpStruct(dst []byte, v interface{}) []byte {
	srcV := reflect.ValueOf(v)

	srcP := unsafe.Pointer(srcV.Pointer())

	size := srcV.Elem().Type().Size()
	if len(dst) >= int(size) {
		dst = dst[:size]
	} else {
		dst = make([]byte, size)
	}
	BufH := (*reflect.SliceHeader)(unsafe.Pointer(&dst))

	MemMove(unsafe.Pointer(BufH.Data), srcP, size)
	return dst
}
