package main

import (
	"fmt"
	"os"
)

func main() {
	//path to prog
	argsWithProg := os.Args

	//holds actual args
	argsWithoutProg := os.Args[1:]

	//getting individual args by indexing
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
