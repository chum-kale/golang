//use chan as func param
//specify if it is can send or recieve values

package main

import "fmt"

//ping only accepts values
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//The pong function accepts one channel for receives (pings) and a second for sends (pongs).
func pong(pings <-chan string, pongs<- string) {
	msg := <- pings
	pongs <- msg
}

func main() {
	pings := mak(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

//chan<- means the channel can only be sent to (send-only).
//<-chan means the channel can only be received from (receive-only
//no arrows (chan typoe) will allow both sending and receving
//space means passing a value