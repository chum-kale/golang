package main

import (
	"context"
	"fmt"
	"time"
)

func performTask(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second): // Simulate long task
		fmt.Println("Task completed successfully")
	case <-ctx.Done():
		fmt.Println("Task timed out:", ctx.Err()) // Handle timeout
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Ensure resources are released

	go performTask(ctx)

	time.Sleep(4 * time.Second) // Let the main function wait long enough to see the result
}
