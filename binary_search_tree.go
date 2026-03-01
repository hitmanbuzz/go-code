//go:build ignore

// This is a Binary Search Tree (not Binary Tree)
// Nodes are sorted when inserted to the Tree

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
	size uint
}

func NewNode(item int) *Node {
	return &Node{
		data: item,
	}
}

func NewBinaryTree(data int) *BinaryTree {
	return &BinaryTree{
		root: &Node{
			data:  data,
			left:  nil,
			right: nil,
		},
		size: 1,
	}
}

func AddNode(node *Node, item *Node) {
	if item.data < node.data {
		if node.left == nil {
			node.left = item
			return
		} else {
			AddNode(node.left, item)
		}
	} else {
		if node.right == nil {
			node.right = item
			return
		} else {
			AddNode(node.right, item)
		}
	}
}

func SearchNode(node *Node, item int) bool {
	if node == nil {
		return false
	}

	if node.data == item {
		return true
	}

	if item < node.data {
		return SearchNode(node.left, item)
	} else {
		return SearchNode(node.right, item)
	}
}

// Find the smallest node for the given Node
func FindSmallest(node *Node) *Node {
	var smallest *Node

	if node.left != nil {
		return FindSmallest(node.left)
	} else {
		smallest = node
		node = nil
		return smallest
	}
}

func DeleteNode(node *Node, item int) *Node {
	if node == nil {
		return nil
	}

	if item < node.data {
		node.left = DeleteNode(node.left, item)
	} else if item > node.data {
		node.right = DeleteNode(node.right, item)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		successor := FindSmallest(node.right)
		node.data = successor.data
		node.right = DeleteNode(node.right, successor.data)
	}

	return node
}

func (b *BinaryTree) Add(item *Node) {
	AddNode(b.root, item)
	b.size++
}

func (b *BinaryTree) Search(item int) bool {
	isFound := SearchNode(b.root, item)
	return isFound
}

func (b *BinaryTree) Delete(item int) *Node {
	deletedNode := DeleteNode(b.root, item)
	b.root = deletedNode
	b.size--
	return deletedNode
}

func main() {
	a := NewBinaryTree(10)
	a.Add(&Node{data: 5})
	a.Add(&Node{data: 8})
	a.Add(&Node{data: 20})
	a.Add(&Node{data: 30})
	a.Add(&Node{data: 35})
	a.Add(&Node{data: 25})

	fmt.Println("Orignal Total Nodes:", a.size)

	isFound := a.Search(25)
	fmt.Println("Found:", isFound)

	// the original node has been remove and change depending on the tree structure
	fmt.Println("Original Node:", a.root.right.data)
	deletedNode := a.Delete(20)
	if deletedNode == nil {
		fmt.Println("Node not found, unable to delete")
	} else {
		fmt.Println("Node deleted")
	}
	deletedNode = a.Delete(30)
	if deletedNode == nil {
		fmt.Println("Node not found, unable to delete")
	} else {
		fmt.Println("Node deleted")
	}

	fmt.Println(a.root.right.data) // Output: 35

	/*
		Structure (Before node deletion):

								10
							   /  \
						      5    20
							   \    \
						        8   30
						            / \
						           25  35
	*/

	fmt.Println("After Total Nodes:", a.size)
}
