package libuseful

import (
	"unsafe"
)

// SysMemStat represents a global system statistic that is managed atomically.
//
// This type must structurally be a uint64 so that mstats aligns with MemStats.
type SysMemStat uint64

//go:linkname SysAlloc runtime.sysAlloc
func SysAlloc(n uintptr, sysStat *SysMemStat) unsafe.Pointer

//go:linkname SysFree runtime.sysFree
func SysFree(v unsafe.Pointer, n uintptr, sysStat *SysMemStat)
