package main

import "fmt"

/*

Sorted BST - Minimal Tree

Given a sorted (increasing order) array with uniqur integer elements,
write an algorithm to create a binary search tree with minimal height

*/

type Node struct {
	Item        int
	Left, Right *Node
}

func NewNode(i int) *Node {
	return &Node{Item: i, Left: nil, Right: nil}
}

func main() {
	sortedItems := []int{2, 5, 8, 14, 35, 42, 64, 73, 85, 92}
	root := makeBinarySearchTree(sortedItems)
	fmt.Println("Inorder Traversal")
	inorder(root)
}

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Println(root.Item)
	inorder(root.Right)
}

func makeBinarySearchTree(a []int) *Node {

	if len(a) == 0 {
		return nil
	}

	mid := len(a) / 2
	root := NewNode(a[mid])
	root.Left = makeBinarySearchTree(a[:mid])
	root.Right = makeBinarySearchTree(a[mid+1:])

	return root

}
