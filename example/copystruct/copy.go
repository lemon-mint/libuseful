package main

import (
	"fmt"

	"github.com/lemon-mint/libuseful"
)

type s1 struct {
	Exported1 int
	Exported2 bool

	unExported1 int
	unExported2 bool
}

func main() {
	a := s1{
		Exported1:   1024,
		Exported2:   true,
		unExported1: 2048,
		unExported2: true,
	}
	b := s1{}
	fmt.Println("A:", a)
	fmt.Println("B:", b)
	libuseful.CopyStruct(&b, &a)
	fmt.Println("A:", a)
	fmt.Println("B:", b)
}
