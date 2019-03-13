package main

import "fmt"

type Node struct {
	Item        int
	Left, Right *Node
}

func NewNode(n int) *Node {
	return &Node{Item: n, Left: nil, Right: nil}
}

//TODO
//Binary tree as opposed to BST only criteria is max of 2 nodes.
//no restriction that left is less and right is greater

// first common ancestor - find the common ancestor

// Shoudl we keep links to parents so we can make our way up

// Given two nodes n1 and n2; we have to find if they have a common ancestor

func lca(root *Node, n1, n2 int) *Node {

	if root == nil {
		return nil
	}

	// if one of them matches report that
	if root.Item == n1 || root.Item == n2 {
		return root
	}
	// Look for keys in left and right subtrees
	leftlca := lca(root.Left, n1, n2)
	rightlca := lca(root.Right, n1, n2)

	//if leftlca and rightlca are not null, then the node is in one of the subtrees, this is lca
	if leftlca != nil && rightlca != nil {
		return root
	}

	//if leftlca is no nil, return leftlca
	if leftlca != nil {
		return leftlca
	}
	return rightlca

}

func main() {

	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)
	root.Right.Left = NewNode(6)
	root.Right.Right = NewNode(7)

	testable := []struct {
		n1, n2   int
		expected int
	}{
		{4, 5, 2},
		{4, 6, 1},
		{2, 4, 2},
	}

	for _, i := range testable {
		resultlca := lca(root, i.n1, i.n2)
		if resultlca == nil {
			fmt.Println("nil result")
		} else {
			if resultlca.Item != i.expected {
				fmt.Println("expected:", i.expected, "got:", resultlca)
			} else {
				fmt.Println("Pass")
			}
		}

	}
}

/*
Given an array of duple with (parent, child) relationship of tree(s), find out the node that has one parent and no parent?

Given an array of duple with (parent, child) relationship of tree(s), check if two nodes has command ancestor?

*/
