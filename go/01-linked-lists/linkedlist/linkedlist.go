// Package linkedlist provides a basic LinkedList type and methods to handle it, like Push(), Pop(), etc.
package linkedlist

import "fmt"

// Node is each node in the LinkedList
type Node struct {
	Value int
	Next  *Node
}

// LinkedList has two pointers:
// Head >> points to the first Node.
// Tail >> points to the last Node.
type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

// New creates a new LinkedList and sets its first node value
func New(value int) *LinkedList {
	n := &Node{value, nil}
	ll := &LinkedList{n, n, 1}
	return ll
}

// Push adds a new node to the end of the Linked List
func (ll *LinkedList) Push(value int) *LinkedList {
	n := &Node{value, nil}

	if ll.Head == nil {
		ll.Head = n
		ll.Tail = n
	} else {
		ll.Tail.Next = n
		ll.Tail = n
	}

	ll.Tail.Next = nil
	ll.Length++

	return ll
}

// Print shows the details of the Linked List on the screen
func (ll LinkedList) Print() {
	currentNode := ll.Head

	for currentNode != nil {
		fmt.Print(currentNode.Value, "->")
		currentNode = currentNode.Next
	}

	fmt.Print("nil")

	fmt.Println()

	fmt.Println("Head:", ll.Head.Value)
	fmt.Println("Tail:", ll.Tail.Value)
	fmt.Println("Length:", ll.Length)
}