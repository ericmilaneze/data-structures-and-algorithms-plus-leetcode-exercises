package linkedlist

import (
	"fmt"
	"testing"
)

func Example() {
	ll := New(10)
	ll.Push(11)
	ll.Push(12)
	ll.Push(13)
	ll.Push(14)

	ll.Print()

	// Output:
	// 10->11->12->13->14->nil
	// Head: 10
	// Tail: 14
	// Length: 5
}

func ExampleNew() {
	ll := New(10)

	fmt.Println(ll.Head.Value)
	fmt.Println(ll.Tail.Value)
	fmt.Println(ll.Length)

	// Output:
	// 10
	// 10
	// 1
}

func ExampleLinkedList_Push() {
	ll := New(10)
	ll.Push(11)
	ll.Push(12)
	ll.Push(13)
	ll.Push(14)

	fmt.Println(ll.Head.Value)
	fmt.Println(ll.Tail.Value)
	fmt.Println(ll.Length)

	// Output:
	// 10
	// 14
	// 5
}

func TestCreateNewLinkedList(t *testing.T) {
	v := 10
	length := 1

	ll := New(v)

	ll.assertLength(t, length)
	ll.assertHeadIsDefined(t)
	ll.assertTailIsDefined(t)
	ll.assertHeadValue(t, v)
	ll.assertTailValue(t, v)
	ll.assertTailNextIsNil(t)
}

func TestPushOneItem(t *testing.T) {
	vi := 1
	vp := 2
	length := 2

	ll := New(vi)
	r := ll.Push(vp)

	ll.assertSame(t, r, "Push()")
	ll.assertLength(t, length)
	ll.assertHeadIsDefined(t)
	ll.assertTailIsDefined(t)
	ll.assertHeadValue(t, vi)
	ll.assertTailValue(t, vp)
	ll.assertTailNextIsNil(t)
}

func TestPushOneItemToAnEmptyLinkedList(t *testing.T) {
	vp := 2
	length := 1

	ll := &LinkedList{nil, nil, 0}
	r := ll.Push(vp)

	ll.assertSame(t, r, "Push()")
	ll.assertLength(t, length)
	ll.assertHeadIsDefined(t)
	ll.assertTailIsDefined(t)
	ll.assertHeadValue(t, vp)
	ll.assertTailValue(t, vp)
	ll.assertTailNextIsNil(t)
}

func TestPop(t *testing.T) {
	ll := New(10)
	ll.Push(11)
	ll.Push(12)
	ll.Push(13)
	ll.Push(14)

	n := ll.Pop()

	if n == nil {
		t.Error("it should return the last value of the Linked List")
	}

	if n.Value != 14 {
		t.Error("popped value - want 14, got ", n.Value)
	}

	if n.Next != nil {
		t.Error("next item should be nil")
	}

	if ll.Length != 4 {
		t.Error("list length - want ", 4, ", got ", ll.Length)
	}

	if ll.Tail == nil {
		t.Error("tail should not be nil")
	}

	if ll.Tail.Value != 13 {
		t.Error("tail value - want 13, got ", ll.Tail.Value)
	}
}

func TestPopAllItems(t *testing.T) {
	ll := New(10)
	ll.Push(11)
	ll.Push(12)
	ll.Push(13)
	ll.Push(14)

	ll.Pop()
	ll.Pop()
	ll.Pop()
	ll.Pop()
	ll.Pop()
	n := ll.Pop()

	if n != nil {
		t.Error("it should return nil")
	}

	if ll.Length != 0 {
		t.Error("list should be empty")
	}

	if ll.Tail != nil {
		t.Error("tail should be nil")
	}

	if ll.Head != nil {
		t.Error("head should be nil")
	}
}

func (ll *LinkedList) assertSame(t *testing.T, r *LinkedList, funcName string) {
	if r != ll {
		t.Error("func ", funcName, " is returning the wrong pointer")
	}
}

func (ll LinkedList) assertHeadIsDefined(t *testing.T) {
	if ll.Head == nil {
		t.Error("head wasn't defined")
	}
}

func (ll LinkedList) assertLength(t *testing.T, length int) {
	if ll.Length != length {
		t.Error("want ", length, ", got ", ll.Length)
	}
}

func (ll LinkedList) assertTailNextIsNil(t *testing.T) {
	if ll.Tail.Next != nil {
		t.Error("tail should point to nil")
	}
}

func (ll LinkedList) assertHeadValue(t *testing.T, v int) {
	if ll.Head.Value != v {
		t.Error("want ", v, ", got ", ll.Head.Value)
	}
}

func (ll LinkedList) assertTailValue(t *testing.T, v int) {
	if ll.Tail.Value != v {
		t.Error("want ", v, ", got ", ll.Tail.Value)
	}
}

func (ll LinkedList) assertTailIsDefined(t *testing.T) {
	if ll.Tail == nil {
		t.Error("tail wasn't defined")
	}
}
