//atomic counters - comm. bw channels
//same use case as worker pools

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//atomic integer type
	var ops atomic.Uint64

	//waitgroup to wait for all routines to finish their work
	var wg sync.WaitGroup

	//50 goroutines that each increment the counter 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		//increment go counter
		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}
			wg.Done()
		}()
	}

	//wait for completion of goroutines
	wg.Wait()
	fmt.Println("ops:", ops.Load())
}
