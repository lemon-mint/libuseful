package main

import (
	"fmt"
	"unsafe"

	"github.com/lemon-mint/libuseful"
)

type s1 struct {
	Exported1 int  `json:"aaa"`
	Exported2 bool `json:"bbb"`

	unExported1 int
	unExported2 bool
}

func main() {
	a := s1{
		Exported1:   12345678,
		Exported2:   true,
		unExported1: 87654321,
		unExported2: true,
	}
	var buf = make([]byte, int(unsafe.Sizeof(a)))
	fmt.Println(libuseful.DumpStruct(buf, &a))
}
