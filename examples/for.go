package main

import "fmt"

func main() {
	//single condition loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//typical loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	//typical loop with range
	for i := range 3 {
		fmt.Println("range", i)
	}

	//loop without condition
	//runs continously until break or rteurn from func
	for {
		fmt.Println("loop")
		break
	}

	//nested loop
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
