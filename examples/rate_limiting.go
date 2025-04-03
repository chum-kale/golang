//rate limiting - control resource utilization
//works with goroutines, channels, tickers

package main

import (
	"fmt"
	"time"
)

func main() {
	//limit handling incoming requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	//channel that recieves time.time values at specified intervals
	limiter := time.Tick(200 * time.Millisecond)

	//blocking on a recieve from limiter channel - limit to 1 req every 200ms
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//buffering limiter channel
	//allow busrts of up to 3 events
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	//adding new value to busrty limiter
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	//simulating 5 more - 3 will get burst effect
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
