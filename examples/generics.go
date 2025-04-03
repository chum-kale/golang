//generics - type parameters
//generics - functions and data structures that can operate on any type

package main

import "fmt"

// As an example of a generic function, MapKeys takes a map of any type and returns a slice of its keys.
// This function has two type parameters - K and V; K has the comparable constraint, meaning that we can compare values of this type with the == and != operators.
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// example of generic type - List is a singly-linked list with values of any type.
// stacks, queueus are examples as well
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// methods on generic types
// keep the type parameters in place. The type is List[T], not List.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

// invoking generic functions relies on type inference, compiler autos that
func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys:", MapKeys(m))

	_ = MapKeys[int, string](m)
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
