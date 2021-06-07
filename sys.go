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
