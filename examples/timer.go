//timer - for a single event in the future
//it provides a channel that will be notified at the time (2 here)

package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	//block timer channel c until a value is sent idicating timer fired
	<-timer1.C
	fmt.Println("Timer 1 fired")

	//cancel the timer before it fires
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	fmt.Println("Timer 2 stopped")
}
