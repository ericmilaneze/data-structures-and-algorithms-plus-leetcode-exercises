package main

import (
	"fmt"

	"github.com/ericmilaneze/data-structures-and-algorithms-plus-leetcode-exercises/go/01-linked-lists/linkedlist"
)

func main() {
	ll := linkedlist.New(9)
	ll.Print()

	fmt.Println()

	ll.Push(10)
	ll.Push(11)
	ll.Push(12)
	ll.Print()
}
