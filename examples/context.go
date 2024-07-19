//context - carries deadlines, cancellation signals, and other request-scoped values across API boundaries and goroutines.

package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	//creating context.context
	ctx := req.context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	//check done channel of context and return as soon as possible
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		//explain why done channel is closed
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
