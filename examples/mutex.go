//muitex - complex state management
//safely access data across multiple goroutines

package main

import (
	"fmt"
	"sync"
)

//counter is a struct mapas we are dealing with data concurrently
//mutex synchronizes access
//should be passed around wity pointer

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

//lock mutex before accessing counters
//unlock using defer statement

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	//the zero value of a mutex is usable as-is, so no initialization is required here.
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	//increment named counters
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	//running several goroutines concurrently
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}
