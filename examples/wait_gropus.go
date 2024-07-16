//wait for multiple goroutines to finish

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	//waitgroup waits for all goroutines launched to finish
	//wiatgrouip can only be passed to functions via pointer
	var wg sync.WaitGroup

	//launch orouthines
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		//closure that tells waitgrp that worker is done
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}
