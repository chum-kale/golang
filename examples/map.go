package main

import (
	"fmt"
	"maps"
)

func main() {
	//map syntax : make(map[key-type]val-type).
	m := make(map[string]int)

	//set key value pairs
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	//print value via keys
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	//null returns zero
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	//delete
	delete(m, "k2")
	fmt.Println("map:", m)

	//clear empties the map
	clear(m)
	fmt.Println("map:", m)

	//second value after _ indicates wether the key was present in the map or not
	//this means zero values like 0 or ""
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//declare and initialize new map in same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	//utility functions
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
