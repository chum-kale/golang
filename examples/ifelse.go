package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	//using logical operators
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 7 or 8 is even")
	}

	//staement can precede conditionals
	//declared variables are available ina all current and following branches
	if num := 9; num < 0 {
		fmt.Println("nume is negative")
	} else if num < 10 {
		fmt.Println("num has 1 digit")
	} else {
		fmt.Println("num has multiple digit")
	}
}

//go has no ternry operator so use if everywhere
