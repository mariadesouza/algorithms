package binarysearchtree

import "fmt"

//Node: Node of a Binary Search tree
type Node struct {
	Data        int
	Left, Right *Node
}

func NewNode(d int) *Node {
	n := Node{Data: d, Left: nil, Right: nil}
	return &n
}

// Space O(n)

// BinarySearchTree -
// left < node < right
// no duplicate nodes
// O(logn)
func Insert(n *Node, d int) error {
	if n == nil {
		n = NewNode(d)
		return nil
	}
	if d < n.Data {
		if n.Left == nil {
			fmt.Println(d, "is Left of ", n.Data)
			n.Left = NewNode(d)
		} else {
			Insert(n.Left, d)
		}
	} else if d > n.Data {
		if n.Right == nil {
			fmt.Println(d, "is Right of ", n.Data)
			n.Right = NewNode(d)
		} else {
			Insert(n.Right, d)
		}
	} else { // equal - duplicate errror
		return fmt.Errorf("Duplication")
	}

	return nil
}

// InOrder Traversal  => left - node- right
func InOrder(n *Node) {
	if n != nil {
		InOrder(n.Left)
		fmt.Println(n.Data)
		InOrder(n.Right)
	}
}

// PreOrder Traversal => node- left- right
// node before child
// root will always be the first node
// this is depth first search - visit a node and then its neighbors.
func PreOrder(n *Node) {
	if n != nil {
		fmt.Println(n.Data)
		PreOrder(n.Left)
		PreOrder(n.Right)
	}
}

// PostOrder Traversal => left- right- node
// child before node
// root will always be the last node
func PostOrder(n *Node) {
	if n != nil {

		PostOrder(n.Left)
		PostOrder(n.Right)
		fmt.Println(n.Data)
	}
}

// Search O(logn)
func Search(n *Node, d int) *Node {
	if n != nil {
		if d == n.Data {
			return n
		} else if d < n.Data {
			return Search(n.Left, d)
		} else if d > n.Data {
			return Search(n.Right, d)
		}
	}
	return nil
}
