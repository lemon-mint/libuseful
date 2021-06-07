package libuseful

import (
	"unsafe"
)

// MemMove copies n bytes from "from" to "to".
//
// MemMove ensures that any pointer in "from" is written to "to" with
// an indivisible write, so that racy reads cannot observe a
// half-written pointer. This is necessary to prevent the garbage
// collector from observing invalid pointers, and differs from MemMove
// in unmanaged languages. However, MemMove is only required to do
// this if "from" and "to" may contain pointers, which can only be the
// case if "from", "to", and "n" are all be word-aligned.
//
// Implementations are in memmove_*.s.
func MemMove(to, from unsafe.Pointer, n uintptr)

//MemEqual compares a and b
func MemEqual(a, b unsafe.Pointer, size uintptr) bool

// Add adds x to pointer p
func Add(p unsafe.Pointer, x uintptr) unsafe.Pointer

// SystemStack runs fn on a system stack.
// If SystemStack is called from the per-OS-thread (g0) stack, or
// if SystemStack is called from the signal handling (gsignal) stack,
// SystemStack calls fn directly and returns.
// Otherwise, SystemStack is being called from the limited stack
// of an ordinary goroutine. In this case, SystemStack switches
// to the per-OS-thread stack, calls fn, and switches back.
// It is common to use a func literal as the argument, in order
// to share inputs and outputs with the code around the call
// to system stack:
//
//	... set up y ...
//	SystemStack(func() {
//		x = bigcall(y)
//	})
//	... use x ...
//
func SystemStack(fn func())

//go:linkname MemMove runtime.memmove
//go:linkname MemEqual runtime.memequal
//go:linkname Add runtime.add
//go:linkname SystemStack runtime.systemstack
