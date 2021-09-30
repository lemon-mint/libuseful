package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/lemon-mint/libuseful"
)

func main() {
	fmt.Println(libuseful.Getpid())
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(libuseful.Getpid())
			for i := 0; i < 1000; i++ {
				runtime.Gosched()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
