package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// Task represents a job to be processed
type Task struct {
	ID       int
	Data     string
	Duration time.Duration
}

// Result represents the output of processing a task
type Result struct {
	TaskID  int
	Output  string
	Success bool
	Error   error
}

// Worker Pool with fixed number of workers
type WorkerPool struct {
	numWorkers int
	tasks      chan Task
	results    chan Result
	done       chan struct{}
	ctx        context.Context
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(ctx context.Context, numWorkers int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		tasks:      make(chan Task, numWorkers),
		results:    make(chan Result, numWorkers),
		done:       make(chan struct{}),
		ctx:        ctx,
	}
}

// Start launches the worker pool
func (p *WorkerPool) Start() {
	// Launch workers
	var wg sync.WaitGroup
	for i := 1; i <= p.numWorkers; i++ {
		wg.Add(1)
		go p.worker(i, &wg)
	}

	// Wait for all workers to complete in a separate goroutine
	go func() {
		wg.Wait()
		close(p.results)
		close(p.done)
	}()
}

// worker processes tasks
func (p *WorkerPool) worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("Worker %d starting\n", id)

	for {
		select {
		case task, ok := <-p.tasks:
			if !ok {
				log.Printf("Worker %d shutting down\n", id)
				return
			}

			// Process the task
			result := p.processTask(id, task)

			// Send result back
			select {
			case p.results <- result:
			case <-p.ctx.Done():
				return
			}

		case <-p.ctx.Done():
			log.Printf("Worker %d cancelled\n", id)
			return
		}
	}
}

// processTask simulates task processing
func (p *WorkerPool) processTask(workerID int, task Task) Result {
	log.Printf("Worker %d processing task %d\n", workerID, task.ID)

	// Simulate work
	time.Sleep(task.Duration)

	return Result{
		TaskID:  task.ID,
		Output:  fmt.Sprintf("Processed task %d with data: %s by worker %d", task.ID, task.Data, workerID),
		Success: true,
	}
}

// SubmitTask adds a task to the pool
func (p *WorkerPool) SubmitTask(task Task) error {
	select {
	case p.tasks <- task:
		return nil
	case <-p.ctx.Done():
		return fmt.Errorf("worker pool is shutdown")
	}
}

// Results returns the results channel
func (p *WorkerPool) Results() <-chan Result {
	return p.results
}

// Shutdown gracefully shuts down the worker pool
func (p *WorkerPool) Shutdown() {
	close(p.tasks)
	<-p.done
}

func main() {
	// Create a context with cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create and start worker pool with 3 workers
	pool := NewWorkerPool(ctx, 3)
	pool.Start()

	// Submit tasks
	numTasks := 10
	go func() {
		for i := 1; i <= numTasks; i++ {
			task := Task{
				ID:       i,
				Data:     fmt.Sprintf("Task data %d", i),
				Duration: time.Duration(i%3+1) * time.Second,
			}

			if err := pool.SubmitTask(task); err != nil {
				log.Printf("Failed to submit task: %v", err)
				return
			}
		}

		// Close tasks channel after all tasks are submitted
		pool.Shutdown()
	}()

	// Collect results
	completedTasks := 0
	for result := range pool.Results() {
		if result.Success {
			log.Printf("Result: %s\n", result.Output)
		} else {
			log.Printf("Task %d failed: %v\n", result.TaskID, result.Error)
		}

		completedTasks++
		if completedTasks == numTasks {
			break
		}
	}

	log.Println("All tasks completed")
}
