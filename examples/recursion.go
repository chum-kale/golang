package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	//closures can also be recursive
	//for this closure needs to be declared with typed var explicitly
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}
