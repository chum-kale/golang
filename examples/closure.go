//closures are anonymous fuctions akin to inline functions

package main

import "fmt"

// return another function which returns int
// returned function closes over variable i to form a closure
func int_seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	//assign func to variable
	next_int := int_seq()

	//variable is updated each time we call the closure
	fmt.Println(next_int())
	fmt.Println(next_int())
	fmt.Println(next_int())

	//new variable assignment resets the value
	newInts := int_seq()
	fmt.Println(newInts())
}
