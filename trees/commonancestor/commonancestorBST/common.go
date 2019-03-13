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

//if  n1 < n < n2  => LCA

//  if n > n1  && n > n2 => traverse left
// if n < n1 && n < n2 => traverse right
// else root is LCA

// if nodes not found in tree return nil

func lca(root *Node, n1, n2 int) *Node {
	if root == nil {
		return nil
	}

	if root.Item == n1 || root.Item == n2 {
		return root
	}

	var leftlca, rightlca *Node

	if root.Item > n1 && root.Item > n2 {
		leftlca = lca(root.Left, n1, n2)
	}

	if root.Item < n1 && root.Item < n2 {
		rightlca = lca(root.Right, n1, n2)
	}

	if leftlca != nil && rightlca != nil {
		return root
	}

	return nil
}

func Insert(root *Node, d int) {

	if root == nil {
		return
	}

	if root.Item <= d {
		if root.Left == nil {
			root.Left = NewNode(d)
		} else {
			Insert(root.Left, d)
		}
	} else {
		if root.Right == nil {
			root.Right = NewNode(d)
		} else {
			Insert(root.Right, d)
		}
	}

}

//left-> node -> right

func Inorder(root *Node) {
	if root != nil {
		Inorder(root.Left)
		fmt.Println("*", root.Item)
		Inorder(root.Right)
	}
}

func CreateBST(a []int) *Node {

	root := NewNode(a[0])

	for _, n := range a[1:] {
		Insert(root, n)
	}
	return root
}

func main() {
	a := []int{20, 8, 22, 4, 12, 10, 14}
	root := CreateBST(a)

	Inorder(root)

	l := lca(root, 4, 14)
	if l != nil {
		fmt.Println("ancestor of 4, 12 is ", l.Item)
	}
	l = lca(root, 23, 34)
	fmt.Println("ancestor of 23, 34 is ", l.Item)
}
