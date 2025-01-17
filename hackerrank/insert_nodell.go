package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

func printSinglyLinkedList(node *SinglyLinkedListNode, sep string, writer *bufio.Writer) {
	for node != nil {
		fmt.Fprintf(writer, "%d", node.data)

		node = node.next

		if node != nil {
			fmt.Fprintf(writer, sep)
		}
	}
}

/*
 * Complete the 'insertNodeAtPosition' function below.
 *
 * The function is expected to return an INTEGER_SINGLY_LINKED_LIST.
 * The function accepts following parameters:
 *  1. INTEGER_SINGLY_LINKED_LIST llist
 *  2. INTEGER data
 *  3. INTEGER position
 */

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	// Write your code here
	if position < 0 {
		return llist
	}

	new_node := &SinglyLinkedListNode{data: data, next: nil}

	if position == 0 {
		new_node.next = llist
		return new_node
	}

	curr := llist
	for i := 0; i < int(position)-1 && curr != nil; i++ {
		curr = curr.next
	}

	new_node.next = curr.next
	curr.next = new_node

	return llist
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	llistCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	llist := SinglyLinkedList{}
	for i := 0; i < int(llistCount); i++ {
		llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		llistItem := int32(llistItemTemp)
		llist.insertNodeIntoSinglyLinkedList(llistItem)
	}

	dataTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	data := int32(dataTemp)

	positionTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	position := int32(positionTemp)

	llist_head := insertNodeAtPosition(llist.head, data, position)func isBalanced(s string) string {
		// Write your code here
		stack := []rune{}
			bracketPairs := map[rune]rune{
			')': '(',
			'}': '{',
			']': '[',
		}
		for _, char := range s {
			switch char {
				case '(', '{', '[':
				stack = append(stack, char)
				case ']', '}', ')':
				if len(stack) == 0 || stack[len(stack) - 1] != bracketPairs[char] {
					return "NO"
				}
				stack = stack[:len(stack)-1]
			}
		}
		
		if len(stack) == 0 {
			return "YES"
		}
		
		return ""
	}

	printSinglyLinkedList(llist_head, " ", writer)
	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
