package main

import "fmt"

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

func BreadthFirstSearch(r *Node) {
	var q []*Node

	//enqueue - append to queue
	q = append(q, r)
	r.Visited = true

	for len(q) > 0 {
		front := q[0]
		fmt.Println(string(front.Item))
		q = q[1:] //dequeue
		for _, c := range front.Children {
			if c.Visited == true {
				continue
			}
			q = append(q, c) //append to queue all children not visited
		}
	}

}

func ResetVisited(r *Node) {
	if r == nil {
		return
	}
	r.Visited = false
	for _, c := range r.Children {
		ResetVisited(c)
	}
}

func main() {
	root := NewNode('A')
	nodeB := NewNode('B')
	nodeC := NewNode('C')
	root.AddChild(nodeB)
	root.AddChild(nodeC)

	nodeD := NewNode('D')
	nodeB.AddChild(nodeD)
	nodeE := NewNode('E')
	nodeB.AddChild(nodeE)

	nodeF := NewNode('F')
	nodeC.AddChild(nodeF)

	BreadthFirstSearch(root)
}
