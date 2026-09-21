//go:build ignore

package main

import "fmt"

// Binary Tree which insert new node to the smallest branch (left or right)

type Node[T any] struct {
	data  T
	left  *Node[T]
	right *Node[T]

	left_size  uint
	right_size uint
}

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{
		data:       data,
		left:       nil,
		right:      nil,
		left_size:  0,
		right_size: 0,
	}
}

type BinaryTree[T any] struct {
	root *Node[T]
}

func NewBinaryTree[T any](data T) *BinaryTree[T] {
	return &BinaryTree[T]{root: NewNode(data)}
}

func insert[T any](node *Node[T], data T) {
	if node.left_size > node.right_size {
		if node.right == nil {
			node.right = NewNode(data)
		} else {
			insert(node.right, data)
		}

		node.right_size++
	} else {
		if node.left == nil {
			node.left = NewNode(data)
		} else {
			insert(node.left, data)
		}

		node.left_size++
	}
}

func (bt *BinaryTree[T]) insert_node(data T) {
	insert(bt.root, data)
}

func main() {
	bt := NewBinaryTree(1)
	bt.insert_node(2)
	bt.insert_node(3)
	bt.insert_node(4)
	bt.insert_node(5)
	bt.insert_node(6)
	bt.insert_node(7)

	fmt.Println("Root:", bt.root.data)
	fmt.Println("Root Left:", bt.root.left.data)
	fmt.Println("Root Right:", bt.root.right.data)

	fmt.Println("\n-------------------------")
	fmt.Println()

	// root left children
	fmt.Println("[ROOT LEFT CHILDREN]")
	fmt.Println("Root Left Left:", bt.root.left.left.data)
	fmt.Println("Root Left Right:", bt.root.left.right.data)

	fmt.Println("\n-------------------------")
	fmt.Println()

	// root right children
	fmt.Println("[ROOT RIGHT CHILDREN]")
	fmt.Println("Root Right Left:", bt.root.right.left.data)
	fmt.Println("Root Right Right:", bt.root.right.right.data)

	fmt.Println(`
      1
    /   \
   2      3
 /  \    / \
4   6   5   7
		`)

	/*

				  1
			   /     \
		      2       3
		     / \     / \
		    4   6   5   7
	*/
}
