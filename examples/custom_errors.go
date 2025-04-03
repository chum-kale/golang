//custom types can be errors by implementing the Error() method on them

package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg     int
	message string
}

// Adding this Error method makes argError implement the error interface
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with this"}
	}
	return arg + 3, nil
}

//errors.as() - checks that a given error (or any error in its chain) matches a
//specific error type and converts to a value of that type, returning
//true. If there’s no match, it returns false.

func main() {
	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
