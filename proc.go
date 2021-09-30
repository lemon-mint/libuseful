package libuseful

import (
	_ "unsafe"
)

//go:linkname ProcPin runtime.procPin
//go:linkname ProcUnpin runtime.procUnpin

func ProcPin() int
func ProcUnpin()

func Getpid() int {
	id := ProcPin()
	ProcUnpin()
	return id
}
