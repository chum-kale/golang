package main

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (list *LinkedList) InsertAtFront(data int) {
	if list.head == nil {
		newNode := &Node{data: data, next: nil}
		list.head = newNode
		return
	}

	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

func (list *LinkedList) InsertInBetween(data int, position int) {
	if position < 0 {
		return
	}

	newNode := &Node{data: data, next: nil}
	if position == 0 {
		newNode.next = list.head
		list.head = newNode
		return
	}

	current := list.head
	for i := 0; i < position-1 && current != nil; i++ {
		current = current.next
	}

	if current == nil {
		return
	}
	newNode.next = current.next
	current.next = newNode
}

func (list *LinkedList) DeleteAtFront() {
	if list.head == nil {
		return
	}
	list.head = list.head.next
	return
}

func (list *LinkedList) DeleteAtEnd() {
	if list.head == nil {
		return
	}

	if list.head.next == nil {
		list.head = nil
		return
	}

	current := list.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
}

func (list *LinkedList) DeleteInBetween(position int) {
	if position < 0 || list.head == nil {
		return
	}

	if position == 0 {
		list.head = list.head.next
		return
	}

	current := list.head
	for i := 0; i < position-1 && current != nil; i++ {
		current = current.next
	}

	if current == nil || current.next == nil {
		return
	}

	current.next = current.next.next
}

func (list *LinkedList) PrintList() {
	current := list.head
	for current != nil {
		print(current.data, " ")
		current = current.next
	}
	println()
}

func (list *LinkedList) Reverse() {
	var prev *Node
	curr := list.head
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	list.head = prev
}
