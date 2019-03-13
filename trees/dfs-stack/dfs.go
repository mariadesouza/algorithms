package main

import (
	"fmt"
)

type Node struct {
	Item     rune
	Visited  bool
	Children []*Node
}

func NewNode(element rune) *Node {
	return &Node{Item: element, Visited: false, Children: nil}
}

func (n *Node) AddChild(c *Node) {
	n.Children = append(n.Children, c)
}

func (n *Node) AddChildElement(element rune) *Node {
	c := NewNode(element)
	n.Children = append(n.Children, c)
	return c
}

/*
   A
  / \
  B  C
  /\  /
  D E F

 A->B->D->E_>C->F

*/
func main() {
	root := NewNode('A')
	nodeB := root.AddChildElement('B')
	nodeC := root.AddChildElement('C')

	nodeB.AddChildElement('D')
	nodeB.AddChildElement('E')
	nodeC.AddChildElement('F')

	BreadthFirstSearch(root)
}

func BreadthFirstSearch(r *Node) {

	var stack []*Node

	// push -append
	// pop - a[:len(s)-1]

	stack = append(stack, r)

	for len(stack) > 0 {

		n := stack[len(stack)-1]
		n.Visited = true
		fmt.Println(string(n.Item))
		stack = stack[:len(stack)-1]

		for _, c := range n.Children {
			if c.Visited == true {
				continue
			}
			stack = append(stack, c)
		}
	}

}
