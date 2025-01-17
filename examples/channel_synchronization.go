//synchronize exexution across goroutines
//example of using a blocking receive to wait for a goroutine to finish.
//When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.

package main

import (
	"fmt"
	"time"
)

// funcrion in goroutine- infomrm another goroutine that functions work is done
func worker(done chan bool) {
	fmt.Print("working....")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	//block till we recieve notif from worker on the channel
	<-done
}
