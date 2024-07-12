//methods - functions defined on struct types

package main

import "fmt"

type rect struct {
	width, height int
}

// reciver of type rect
func (r *rect) area() int {
	return r.width * r.height
}

// can be defined without pointer as well
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
