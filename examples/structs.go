package main

import "fmt"

type person struct {
	name string
	age  int
}

// we can return a pointer as a local variable will survive scope of function
// pointers arent compulsory in structs
func new_person(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(new_person("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	//pointers are automatically derefrenced
	sp := &s
	fmt.Println(sp.age)

	//structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

	//anonymous truct- used for a single value
	//used in table driven tests
	dog := struct {
		name   string
		isGood bool
	}{
		"rex",
		true,
	}
	fmt.Println(dog)
}
