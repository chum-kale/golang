package main

import (
	"fmt"
	"net/http"
)

//handler
//generally use http.HandlerFunc
//http.ResponseWriter and a http.Request as arguments.
//The response writer is used to fill in the HTTP response.

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Printf(w, "hello\n")
}

// this handler reads all http reqs and echoes them into the body
func header(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	//register handler on server routes
	//takes function as an arg
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	//set up router
	http.ListenAndServe(":8090", nil)
}
