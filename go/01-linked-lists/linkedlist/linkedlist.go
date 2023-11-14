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

// Pop removes the tail node and returns it
func (ll *LinkedList) Pop() *Node {
	if ll.Length == 0 {
		return nil
	}

	if ll.Length == 1 {
		t := ll.Head
		ll.Head = nil
		ll.Tail = nil
		ll.Length = 0
		return t
	}

	prev := ll.Head
	curr := ll.Head.Next
	for curr.Next != nil {
		prev = curr
		curr = curr.Next
	}

	ll.Tail = prev
	ll.Tail.Next = nil
	curr.Next = nil
	ll.Length--

	return curr
}

// Unshift adds a new item to the beginning of the Linked List
func (ll *LinkedList) Unshift(value int) *LinkedList {
	temp := ll.Head
	ll.Head = &Node{value, temp}

	if ll.Length == 0 {
		ll.Tail = ll.Head
	}

	ll.Length++

	return ll
}

// Shift removes an items from the beginning of the Linked List
func (ll *LinkedList) Shift() *Node {
	if ll.Length == 0 {
		return nil
	}

	oldHead := ll.Head
	ll.Head = oldHead.Next
	oldHead.Next = nil
	ll.Length--

	if ll.Length == 0 {
		ll.Tail = nil
	}

	return oldHead
}

func (ll LinkedList) Get(pos int) *Node {
	if pos < 0 || pos > ll.Length {
		return nil
	}

	if pos == 0 {
		return ll.Head
	}

	if pos == ll.Length-1 {
		return ll.Tail
	}

	curr := ll.Head.Next
	for i := 1; i < pos; i++ {
		curr = curr.Next
	}

	return curr
}

// Print shows the details of the Linked List on the screen
func (ll LinkedList) Print() {
	if ll.Length > 0 {
		currentNode := ll.Head

		for currentNode != nil {
			fmt.Print(currentNode.Value, "->")
			currentNode = currentNode.Next
		}

		fmt.Print("nil")

		fmt.Println()
		fmt.Println("Head:", ll.Head.Value)
		fmt.Println("Tail:", ll.Tail.Value)
	}

	fmt.Println("Length:", ll.Length)
}
