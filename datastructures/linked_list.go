package datastructures

import "fmt"

/*
	In computer science, a linked list is a linear collection of data elements,
	in which linear order is not given by their physical placement in memory.
	Each pointing to the next node by means of a pointer.
	It is a data structure consisting of a group of nodes which together represent a sequence.
	Here is source code of the Go Program to Implement Single Unsorted Linked List
*/

// Here we create a strucutre to store the data, and references to the prev and next nodes in the list
type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

// This is the reference to our List's head (first), and tail (last)
type List struct {
	head *Node
	tail *Node
}

// We can insert data to the list by also passing references to the Node
func (L *List) ListInsert(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

// Display the specified Key in the Node for the List
func (l *List) ListDisplay() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

// Display the specified Key in the Node
func ListDisplay(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

// Show the previous key in the Node
func ListShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-", list.key)
		list = list.prev
	}
	fmt.Println()
}

// We can reverse the list by switching the references in the node
func (l *List) ListReverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	ListDisplay(l.head)
}

// Return the new list interface
func NewList() List {
	link := List{}
	return link
}

func RunListFunctions() {
	// Create a list
	link := List{}
	// Insert linked data
	link.ListInsert(5)
	link.ListInsert(9)
	link.ListInsert(13)
	link.ListInsert(22)
	link.ListInsert(28)
	link.ListInsert(36)

	fmt.Println("\n==============================")
	// Preview the head and tail of the list
	fmt.Printf("Head: %v\n", link.head.key)
	fmt.Printf("Tail: %v\n", link.tail.key)
	// Display the list with it's link references
	link.ListDisplay()
	fmt.Println("\n==============================")
	// Then we can reverse the list, and it's links
	link.ListReverse()
	// And now we have a new head and tail reference
	fmt.Printf("head: %v\n", link.head.key)
	fmt.Printf("tail: %v\n", link.tail.key)
	fmt.Println("\n==============================")
}
