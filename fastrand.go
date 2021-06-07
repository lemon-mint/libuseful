package libuseful

import (
	_ "unsafe"
)

// FastRand32 implements the xorshift64+ algorithm
func FastRand32() uint32

// FastRand32n is similar to FastRand32() % n, but faster.
func FastRand32n(n uint32) uint32

//go:linkname FastRand32 runtime.fastrand
//go:linkname FastRand32n runtime.fastrandn
