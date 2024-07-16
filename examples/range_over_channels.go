package main

import "fmt"

func main() {
	//range iteration over channel
	//we will use queue as channel
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	//ranged loop
	for elem := range queue {
		fmt.Println(elem)
	}
}
