//handling signals with channels

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//create buffered channel
	//signal notification works by sending os.Signal values on a channel
	sigs := make(chan os.Signal, 1)

	//register given channel to recieve notifs
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	//scenario of graceful shutdown
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	//wait here till it gets expected value on done channel then exit
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
