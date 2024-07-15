package main

//ubuffered - only accepts send (chan <-) if there is a corresponding recieve (chan (<-)

import "fmt"

func main() {
	//channel for buffering 2 values
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
