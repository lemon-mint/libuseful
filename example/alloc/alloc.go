package main

import (
	"fmt"
	"unsafe"

	"github.com/lemon-mint/libuseful"
)

func main() {
	var a int64 = 2048
	var AllocStat libuseful.SysMemStat

	fmt.Println("AllocStat:", AllocStat)

	ptr := libuseful.SysAlloc(8, &AllocStat)

	fmt.Println("AllocStat:", AllocStat)
	fmt.Println(*(*int)(ptr))
	libuseful.MemMove(ptr, unsafe.Pointer(&a), 8)
	fmt.Println(*(*int)(ptr))

	libuseful.SysFree(ptr, 8, &AllocStat)

	fmt.Println("AllocStat:", AllocStat)

}
