package libuseful

import (
	_ "runtime"
	"unsafe"
)

//go:linkname MemMove runtime.memmove
func MemMove(to, from unsafe.Pointer, n uintptr)
