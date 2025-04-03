package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	select {
	case <-time.After(10 * time.Second): // Simulate long task
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Task cancelled:", ctx.Err()) // Handle cancellation
	}
}

func main() {
	deadline := time.Now().Add(2 * time.Second) // Set deadline 2 seconds from now
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go longRunningTask(ctx)

	time.Sleep(5 * time.Second) // Let the main function wait for the task or deadline
}
