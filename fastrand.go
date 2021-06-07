package libuseful

import (
	_ "unsafe"
)

//go:linkname FastRand32 runtime.fastrand
func FastRand32() uint32

//go:linkname FastRand32n runtime.fastrandn
func FastRand32n(n uint32) uint32
