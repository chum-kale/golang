package main

import (
	"context"
	"fmt"
	"time"
)

// longRunningTask listens to the context's Done channel to stop when canceled.
func longRunningTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task was canceled:", ctx.Err())
			return
		default:
			fmt.Println("Task is running...")
			time.Sleep(500 * time.Millisecond) // Simulate work
		}
	}
}

func main() {
	// Create a cancelable context
	ctx, cancel := context.WithCancel(context.Background())

	// Start a long-running task in a goroutine
	go longRunningTask(ctx)

	// Simulate some work in the main function
	time.Sleep(2 * time.Second) // Let the task run for 2 seconds

	// Cancel the context
	fmt.Println("Canceling the task...")
	cancel()

	// Wait a little to see the goroutine's output after cancellation
	time.Sleep(1 * time.Second)
}
