func mergeListsIterative(head1, head2 *SinglyLinkedListNode) *SinglyLinkedListNode {
	// Dummy node to simplify handling of head
	dummy := &SinglyLinkedListNode{}
	current := dummy

	// Traverse both lists and merge
	for head1 != nil && head2 != nil {
		if head1.data <= head2.data {
			current.next = head1
			head1 = head1.next
		} else {
			current.next = head2
			head2 = head2.next
		}
		current = current.next
	}

	// Append the remaining nodes of the non-empty list
	if head1 != nil {
		current.next = head1
	} else {
		current.next = head2
	}

	return dummy.next
}
