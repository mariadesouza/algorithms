package binarytree

// BinaryTree : A binary tree node has data, pointer to left child and a pointer to right child
type Node struct {
	Data        interface{}
	Left, Right *Node
}

func NewNode(d interface{}) *Node {

	n := Node{Data: d, Left: nil, Right: nil}
	return &n
}
