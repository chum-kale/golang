//built in synchronization of goroutines and channels

package main

import (
	"sync/atomic"
	"time"
)

//a single go routine owns the state- master goroutine
type readOp struct {
	key int
	resp chan int
}

type writeop struct {
	key int
	val int
	resp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	//channels fo read and write requests
	reads := make(chan readOP)
	writes := make(chan writeOp)

	//master goroutine - owns the states
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	//each read
	//1. construct redOP
	//2. send it over reads channel
	//3. recieve result over resp chnnel

	//start 100 goroutines to issue reads o master
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key: rand.Intn(5),
					resp: make(chan int)
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep((time.Millisecond))
			}
		}()
	}

	//start 10 writes
	for w := 0; w < 10; w++ {
        go func() {
            for {
                write := writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

	time.Sleep(time.Second)
	//capure and report op counts
	raeaOpsFinal := atomic.LoadInt64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}