package main

import "fmt"

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

func reverse(t *Node) *Node {
	if t == nil {
		return nil
	}
	tmp := t.Left
	t.Left = reverse(t.Right)
	t.Right = reverse(tmp)
	return t
}

func inorder(t *Node) {
	if t == nil {
		return
	}
	inorder(t.Left)
	fmt.Print(t.Item, " ")
	inorder(t.Right)
}

func main() {
	root := NewNode(5)
	root.AddLeftNode(8)
	r1 := root.AddRightNode(10)
	r1.AddRightNode(3)

	inorder(root)
	fmt.Println()
	t := reverse(root)
	fmt.Println("After Reverse:")
	inorder(t)

}
