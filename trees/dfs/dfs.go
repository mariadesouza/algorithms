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

	DepthFirstSearch(root)
	ResetVisited(root)
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

// start with root and explore all branches completely and then backtrack
func DepthFirstSearch(r *Node) {

	if r == nil {
		return
	}

	if r.Visited == true {
		return
	}
	fmt.Println(string(r.Item))
	r.Visited = true
	for _, child := range r.Children {
		DepthFirstSearch(child)
	}

}
