package main

import "fmt"

type Node struct {
	value     string
	visited   bool
	neighbors []*Node
}

type Graph map[string]*Node

func NewGraph() *Graph {
	g := Graph(make(map[string]*Node))
	return &g
}

func (g *Graph) AddNode(i ...string) {
	if g == nil {
		g = NewGraph()
	}

	for _, str := range i {
		n := Node{value: str, visited: false, neighbors: nil}
		(*g)[str] = &n
	}
}

func (g *Graph) AddEdge(i string, node ...*Node) error {
	n, ok := (*g)[i]
	if !ok {
		return fmt.Errorf("No node found with value %d", i)
	} else {
		n.neighbors = append(n.neighbors, node...)
	}
	return nil
}

func (g *Graph) resetVisited() {
	for _, node := range *g {
		node.visited = false
	}
}

func main() {
	g := NewGraph()
	g.AddNode("A", "B", "C", "D")
	g.AddEdge("A", (*g)["B"], (*g)["C"])
	g.AddEdge("B", (*g)["C"])
	g.AddEdge("C", (*g)["A"], (*g)["D"])
	g.AddEdge("D", (*g)["D"])
}

func isReachable(s string, d string) bool {
	//reset visited
	// do BFS
	return false
}
