package libuseful

import (
	"unsafe"
)

// SysMemStat represents a global system statistic that is managed atomically.
//
// This type must structurally be a uint64 so that mstats aligns with MemStats.
type SysMemStat uint64

// SysAlloc allocates unmanaged memory
func SysAlloc(n uintptr, sysStat *SysMemStat) unsafe.Pointer

// SysFree frees memory allocated by SysAlloc
func SysFree(v unsafe.Pointer, n uintptr, sysStat *SysMemStat)

//go:linkname SysAlloc runtime.sysAlloc
//go:linkname SysFree runtime.sysFree

var libUsefulMemStat SysMemStat

// Alloc allocates unmanaged memory (Calls SysAlloc)
func Alloc(size uintptr) unsafe.Pointer {
	return SysAlloc(size, &libUsefulMemStat)
}

// Free frees memory allocated by Alloc
func Free(ptr unsafe.Pointer, size uintptr) {
	SysFree(ptr, size, &libUsefulMemStat)
}
