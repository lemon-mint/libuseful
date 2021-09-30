package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/lemon-mint/libuseful"
)

func main() {
	fmt.Println(libuseful.GetPID())
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(libuseful.GetPID())
			for i := 0; i < 1000; i++ {
				runtime.Gosched()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
