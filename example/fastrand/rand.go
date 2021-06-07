package main

import (
	"fmt"

	"github.com/lemon-mint/libuseful"
)

func main() {
	fmt.Println("FastRand32() =", libuseful.FastRand32())
	fmt.Println("FastRand32() =", libuseful.FastRand32())
	fmt.Println("FastRand32() =", libuseful.FastRand32())
	fmt.Println("FastRand32() =", libuseful.FastRand32())
	fmt.Println("FastRand32n(256) =", libuseful.FastRand32n(256))
	fmt.Println("FastRand32n(256) =", libuseful.FastRand32n(256))
	fmt.Println("FastRand32n(256) =", libuseful.FastRand32n(256))
	fmt.Println("FastRand32n(256) =", libuseful.FastRand32n(256))
}
