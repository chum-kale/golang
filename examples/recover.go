//recover - stop panic from aborting code and let it continue instead
//eg- a single client connection shows error, so it has to be shut and should not affect the server

package main

import "fmt"

// panic func
func mayPanic() {
	panic("a problem")
}

// recover must always be in a defer function
func main() {
	//return value is he error raised in the call to panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()
	fmt.Println("After mayPanic()")
}
