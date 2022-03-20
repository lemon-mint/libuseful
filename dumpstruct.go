package libuseful

import (
	"reflect"
	"unsafe"
)

// DumpStruct takes a structure instance v as input and copies the memory to dst.
//
// If the size of dst is smaller than the size of v, a new byte slice is allocated.
func DumpStruct[T any](dst []byte, v *T) []byte {
	size := int(unsafe.Sizeof(*v))
	if len(dst) >= int(size) {
		dst = dst[:size]
	} else {
		dst = make([]byte, size)
	}
	BufH := (*reflect.SliceHeader)(unsafe.Pointer(&dst))

	MemMove(unsafe.Pointer(BufH.Data), unsafe.Pointer(v), uintptr(size))
	return dst
}
