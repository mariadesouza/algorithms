package main

import "fmt"

type Node struct {
	Item        int
	Left, Right *Node
}

func NewNode(i int) *Node {
	return &Node{Item: i, Left: nil, Right: nil}
}

/*
Compute max depth - number of nodes along the longest path

Recursively calculate height of left and right subtrees

return => max(left , right) + 1
*/

func findMaxDepth(n *Node) int {
	if n == nil {
		return 0
	}
	lDepth := findMaxDepth(n.Left)
	rDepth := findMaxDepth(n.Right)

	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

func main() {
	b := NewNode(5)
	b.Left = NewNode(8)
	b.Right = NewNode(3)

	b.Left.Left = NewNode(10)
	c := b.Left.Left
	c.Left = NewNode(9)
	c.Right = NewNode(11)

	fmt.Println(findMaxDepth(b))
}
