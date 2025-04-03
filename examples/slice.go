package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	//empty slice with nonzro length
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	//set and get like arrays

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	//len returns the length of the slice as expected.

	fmt.Println("len:", len(s))

	//append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	//copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	//slice operations within a slice (akin to python array)
	//slice[low:high]. For example, this gets a slice of the elements s[2], s[3], and s[4].
	l := s[2:5]
	fmt.Println("sl1:", l)
	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)

	//utility functions - equal
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	//multidimensional slices - length of inner slices can vary unlike arrays
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
