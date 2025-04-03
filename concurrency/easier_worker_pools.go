package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d starting job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // Send result back
	}
}

func main() {
	jobs := make(chan int, 5)    // Channel for jobs
	results := make(chan int, 5) // Channel for results

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // No more jobs will be sent

	// Collect results
	for a := 1; a <= 5; a++ {
		<-results
	}
}
