//sort something other than its natural order

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	//compare string lengths
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	//sort by length
	slices.SortFunc(fruits, lenCmp)

	//sort something that isn't a built in type - structs
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}
