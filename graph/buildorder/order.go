package main

import "fmt"

type Node struct {
	Item  string
	Edges []*Node
}

type Graph []*Node

func (g *Graph) AddVertex(values ...string) {

	for _, s := range values {
		*g = append(*g, &Node{Item: s, Edges: nil})
	}
}

func (g *Graph) FindBuildOrder(n *Node, resolved []string, lookup map[string]struct{}, unresolved map[string]struct{}) []string {
	var e struct{}
	unresolved[n.Item] = e
	for _, e := range n.Edges {
		if _, ok := lookup[e.Item]; !ok { // we jhavent resolved the edge
			if _, seen := unresolved[e.Item]; seen { // circular
				fmt.Println("Circular dependency error")
				return nil
			}
			resolved = g.FindBuildOrder(e, resolved, lookup, unresolved)
		}
	}

	lookup[n.Item] = e
	resolved = append(resolved, n.Item)

	delete(unresolved, n.Item)
	return resolved
}

func (g *Graph) AddEdge(index int, edges ...*Node) {
	for _, e := range edges {
		(*g)[index].Edges = append((*g)[index].Edges, e)
	}
}

func main() {
	var g Graph
	g.AddVertex("A", "B", "C") //0 , 1, 2
	g.AddVertex("D", "E")      //3, 4

	g.AddEdge(0, g[1]) // a depends on b
	g.AddEdge(0, g[3]) // a-> d
	g.AddEdge(1, g[2]) //b -> c
	g.AddEdge(1, g[4]) //b -> e
	g.AddEdge(2, g[3]) //c -> d
	g.AddEdge(2, g[4]) // c-> e

	// circular
	//g.AddEdge(3, g[1])

	var resolved []string
	lookup := make(map[string]struct{})
	unresolved := make(map[string]struct{})
	resolved = g.FindBuildOrder(g[0], resolved, lookup, unresolved)
	fmt.Println("Build Order:")
	for _, s := range resolved {
		fmt.Print(" ", s)
	}
	fmt.Println()
}
