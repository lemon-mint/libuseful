package libuseful

import (
	"reflect"
	"unsafe"
)

// CopyStruct copies struct src to dst
func CopyStruct(dst, src interface{}) {
	dstV := reflect.ValueOf(dst)
	srcV := reflect.ValueOf(src)

	dstP := unsafe.Pointer(dstV.Pointer())
	srcP := unsafe.Pointer(srcV.Pointer())

	MemMove(dstP, srcP, dstV.Elem().Type().Size())
}
