package main

import (
	"fmt"
	"time"
)

// worker recieves work on job channel and sends results to results.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	//start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//send 5 jobs and close the channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	//result collection
	//alternative way  WaitGroup
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
