//timeouts - handle connections with external resources

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	//buffered channel - This is a common pattern to prevent goroutine leaks in case the channel is never read.
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	//timeout of 1 (result1 canned)
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	//timeout 3 - result 2 succeeds
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}
