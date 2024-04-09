package main

import "fmt"

type info struct {
	email string
	zip   int
}

type person struct {
	first_name string
	last_name  string
	contact    info
}

func main() {
	//Jams := person{first_name: "Jams", last_name: "Jake"}
	jim := person{
		first_name: "Jim",
		last_name:  "Jack",
		contact: info{
			email: "jim@gmail.com",
			zip:   79879,
		},
	}
	jim_pointer := &jim
	jim_pointer.update_name("Carl")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointer_to_person *person) update_name(new_name string) {
	(*pointer_to_person).first_name = new_name
}
