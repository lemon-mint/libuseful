package libuseful

import "unsafe"

type SysMemStat uint64

//go:linkname SysAlloc runtime.sysAlloc
func SysAlloc(n uintptr, sysStat *SysMemStat) unsafe.Pointer

//go:linkname SysFree runtime.sysFree
func SysFree(v unsafe.Pointer, n uintptr, sysStat *SysMemStat)
