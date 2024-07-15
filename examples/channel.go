//channels - pipes that connect goroutines
//send values into one 1 goroutine and recieve them in another goroutine

package main

import "fmt"

func main() {
	//channels have type of values they convey
	messages := make(chan string)

	//send value to channel using <- syntax
	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
