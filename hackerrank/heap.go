package main

import (
	"fmt"
)

// MinHeap structure
type MinHeap struct {
	data []int
}

// Insert adds a new element to the heap
func (h *MinHeap) Insert(x int) {
	h.data = append(h.data, x)
	h.heapifyUp(len(h.data) - 1)
}

// Remove removes an element from the heap
func (h *MinHeap) Remove(x int) {
	// Find the index of the element to remove
	index := -1
	for i, val := range h.data {
		if val == x {
			index = i
			break
		}
	}
	if index == -1 {
		return // Element not found
	}

	// Replace with the last element and heapify
	h.data[index] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.heapifyDown(index)
}

// GetMin returns the smallest element in the heap
func (h *MinHeap) GetMin() int {
	if len(h.data) == 0 {
		return -1 // Return -1 for an empty heap
	}
	return h.data[0]
}

// heapifyUp maintains the heap property after insertion
func (h *MinHeap) heapifyUp(index int) {
	parent := (index - 1) / 2
	if index > 0 && h.data[index] < h.data[parent] {
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		h.heapifyUp(parent)
	}
}

// heapifyDown maintains the heap property after removal
func (h *MinHeap) heapifyDown(index int) {
	smallest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < len(h.data) && h.data[left] < h.data[smallest] {
		smallest = left
	}
	if right < len(h.data) && h.data[right] < h.data[smallest] {
		smallest = right
	}

	if smallest != index {
		h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
		h.heapifyDown(smallest)
	}
}

func main() {
	var q int
	fmt.Scan(&q) // Number of queries

	heap := MinHeap{}

	for i := 0; i < q; i++ {
		var queryType int
		fmt.Scan(&queryType)

		switch queryType {
		case 1: // Add element to the heap
			var x int
			fmt.Scan(&x)
			heap.Insert(x)

		case 2: // Remove element from the heap
			var x int
			fmt.Scan(&x)
			heap.Remove(x)

		case 3: // Print the minimum element
			fmt.Println(heap.GetMin())
		}
	}
}
