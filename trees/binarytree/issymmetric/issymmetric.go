package main

import "fmt"

/*
A binary tree is symmetric if left and right nodes are symmetric

*/

type Node struct {
	Item        int
	Left, Right *Node
}

func NewNode(value int) *Node {
	return &Node{Item: value, Left: nil, Right: nil}
}

func (r *Node) AddLeftNode(value int) *Node {
	r.Left = NewNode(value)
	return r.Left
}

func (r *Node) AddRightNode(value int) *Node {
	r.Right = NewNode(value)
	return r.Right
}

// symmetric both nodes are not nil
func isSymetric(n1 *Node, n2 *Node) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 == nil || n2 == nil {
		return false
	}
	return isSymetric(n1.Left, n2.Right)

}

func main() {

	root := NewNode(5)
	root.AddLeftNode(8)
	root.AddRightNode(10)
	//r1.AddRightNode(3)

	fmt.Println(isSymetric(root.Left, root.Right))

}
