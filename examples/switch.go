package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	//normal siwtch case
	fmt.Println("write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//switch statement with multiple expressions
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's the weekday")
	}

	//usage can be similar to if-else loops
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's afternoon")
	}

	//type switch - used to discover type of an interface value
	what_am_i := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Boolean")
		case int:
			fmt.Println("Integer")
		default:
			fmt.Println("Ion know the type :(")
		}
	}
	what_am_i(true)
	what_am_i(1)
	what_am_i("hey")
}
