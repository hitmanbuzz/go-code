//go:build ignore

// Doubly Circular Linked List (DCLL)

package main

import "fmt"

type Node[T comparable] struct {
	Data T
	Next *Node[T]
	Prev *Node[T]
}

func NewNode[T comparable](data T) *Node[T] {
	return &Node[T]{
		Data: data,
		Next: nil,
		Prev: nil,
	}
}

type LinkedList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
	Size uint
}

func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

func (ll *LinkedList[T]) InsertFront(data T) {
	node := NewNode(data)

	if ll.Head == nil {
		node.Next = node
		node.Prev = node

		ll.Head = node
		ll.Tail = node
		ll.Size++
		return
	}

	node.Next = ll.Head
	node.Prev = ll.Tail

	ll.Head.Prev = node
	ll.Tail.Next = node

	ll.Head = node

	ll.Size++
}

func (ll *LinkedList[T]) InsertBack(data T) {
	if ll.Head == nil {
		ll.InsertFront(data)
		return
	}

	node := NewNode(data)
	node.Next = ll.Head
	node.Prev = ll.Tail

	ll.Head.Prev = node
	ll.Tail.Next = node

	ll.Tail = node
	ll.Size++
}

func (ll *LinkedList[T]) DeleteFront() {
	if ll.Head == nil {
		fmt.Println("Linked List is empty")
		return
	}

	if ll.Head == ll.Tail {
		ll.Head = nil
		ll.Tail = nil
		ll.Size--
		return
	}

	newHead := ll.Head.Next
	newHead.Prev = ll.Tail
	ll.Tail.Next = newHead
	ll.Head = newHead

	ll.Size--
}

func (ll *LinkedList[T]) DeleteBack() {
	if ll.Head == nil {
		fmt.Println("Linked List is empty")
		return
	}

	if ll.Head == ll.Tail {
		ll.Head = nil
		ll.Tail = nil
		ll.Size--
		return
	}

	newTail := ll.Tail.Prev
	newTail.Next = ll.Head
	ll.Head.Prev = newTail

	ll.Tail = newTail
	ll.Size--
}

func (ll *LinkedList[T]) DeleteValue(data T) {
	if ll.Head == nil {
		fmt.Println("Linked List is empty")
		return
	}

	if ll.Head.Data == data {
		ll.DeleteFront()
		return
	}

	if ll.Tail.Data == data {
		ll.DeleteBack()
		return
	}

	// next cuz since head is already checked
	curr := ll.Head.Next

	for {
		if curr.Data == data {
			curr.Prev.Next = curr.Next
			curr.Next.Prev = curr.Prev
			ll.Size--
			return
		}

		curr = curr.Next

		if curr == ll.Head {
			break
		}
	}

	fmt.Printf("%v not found in Linked List\n", data)
}

func (ll *LinkedList[T]) InsertAfter(targetData T, data T) {
	if ll.Head == nil {
		fmt.Println("Linked List is empty")
		return
	}

	curr := ll.Head
	for {
		if curr.Data == targetData {
			newNext := NewNode(data)

			newNext.Prev = curr
			newNext.Next = curr.Next

			curr.Next.Prev = newNext
			curr.Next = newNext

			if curr == ll.Tail {
				ll.Tail = newNext
			}

			ll.Size++
			return
		}

		curr = curr.Next

		if curr == ll.Head {
			break
		}
	}

	fmt.Printf("%v not found in Linked List\n", targetData)
}

func (ll *LinkedList[T]) Search(data T) bool {
	if ll.Head == nil {
		return false
	}

	if ll.Head.Data == data {
		return true
	}

	if ll.Tail.Data == data {
		return true
	}

	curr := ll.Head.Next

	for {
		if curr.Data == data {
			return true
		}

		curr = curr.Next

		if curr == ll.Head {
			break
		}
	}

	return false
}

func (ll *LinkedList[T]) Length() uint {
	return ll.Size
}

func (ll *LinkedList[T]) Clear() {
	ll.Head = nil
	ll.Tail = nil
	ll.Size = 0
}

func (ll *LinkedList[T]) PrintForward() {
	if ll.Head == nil {
		fmt.Println("Linked List is empty")
		return
	}

	curr := ll.Head

	for {
		fmt.Print(curr.Data, " <-> ")
		curr = curr.Next

		if curr == ll.Head {
			break
		}
	}

	fmt.Println()
}

func (ll *LinkedList[T]) PrintBackward() {
	if ll.Head == nil {
		fmt.Println("Linked List is empty")
		return
	}

	curr := ll.Tail

	for {
		fmt.Print(curr.Data, " <-> ")
		curr = curr.Prev

		if curr == ll.Tail {
			break
		}
	}

	fmt.Println()
}

func main() {
	ll := New[uint]()
	ll.InsertFront(3)
	ll.InsertFront(2)
	ll.InsertFront(1)

	ll.InsertBack(4)
	ll.InsertFront(0)

	ll.InsertBack(5)
	ll.PrintForward()

	// ll.DeleteBack()
	// ll.PrintForward()

	// ll.DeleteBack()
	// ll.PrintForward()

	// ll.DeleteValue(2)
	// ll.PrintForward()

	// ll.Clear()
	// ll.PrintForward() // nuke it

	ll.InsertAfter(2, 67)
	ll.PrintForward()
	ll.InsertAfter(67, 69)
	ll.PrintForward()

	isFound := ll.Search(67)
	fmt.Println("Is Found:", isFound)
}
