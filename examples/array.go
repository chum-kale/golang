package main

import "fmt"

func main() {
	//array of 5 digits
	var a [5]int
	fmt.Println("emp:", a)

	//set an element
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	//declare and initialize array in one line
	b := [5]int{1,2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//compiler can count no of elements
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcx:", b)

	//: will zero elements
	b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)

	//multidimensional array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0 j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3}
	}
	fmt.Println("2d: ", twoD)
}
