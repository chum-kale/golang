package main

import "fmt"

type queue struct {
	elements []int
}

func (q *queue) Enqueue(data int) {
	q.elements = append(q.elements, data)
}

func (q *queue) Dequeue() (int, error) {
	if len(q.elements) == 0 {
		return 0, fmt.Errorf("empty quwuw")
	}
	front := q.elements[0]
	q.elements = q.elements[1:]
	return front, nil
}

func (q *queue) Peek() (int, error) {
	if len(q.elements) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.elements[0], nil
}

func (q *queue) Isempty() bool {
	return len(q.elements) == 0
}

func (q *queue) Size() int {
	return len(q.elements)
}

func main() {
	queue := &queue{}

	// Enqueue elements
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	fmt.Println("Queue:", queue.elements)

	// Dequeue elements
	val, _ := queue.Dequeue()
	fmt.Println("Dequeued:", val)
	fmt.Println("Queue after dequeue:", queue.elements)

	// Peek at the front
	front, _ := queue.Peek()
	fmt.Println("Front element:", front)

	// Check if the queue is empty
	fmt.Println("Is queue empty?", queue.IsEmpty())

	// Get the size of the queue
	fmt.Println("Size of queue:", queue.Size())
}
