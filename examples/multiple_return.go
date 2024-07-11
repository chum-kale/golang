package main

import "fmt"

// the func takes no args and returns 2 ints
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//_ one value gives a subset of return values
	_, c := vals()
	fmt.Println(c)
}
