//basic sends recieves are blocking
//we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.
//a non-blocking channel operation is an operation that does not cause the goroutine to wait if the channel is not ready for communication.

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	//msg cannot be sent as here is no buffer and reciever
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	//multiway non-blocking se;ect
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no actvity")
	}
}
