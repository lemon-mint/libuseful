package main

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/lemon-mint/libuseful"
)

func main() {
	var a = []byte{
		0, 1, 2, 3,
	}
	var b = []byte{
		0, 0, 0, 0,
	}

	A := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	B := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	fmt.Println("A:", a)
	fmt.Println("B:", b)
	libuseful.MemMove(unsafe.Pointer(B.Data), unsafe.Pointer(A.Data), 4)
	fmt.Println("libuseful.MemMove(B, A, 4)")
	fmt.Println("A:", a)
	fmt.Println("B:", b)
}
