package main

import (
	"fmt"

	"github.com/mariadesouza/algorithms/graph"
)

func main() {

	var g graph.Graph
	g.AddVertex("A", "B", "C")
	g.AddVertex("D", "E")

	g.AddEdge(0, g[1])
	g.AddEdge(1, g[2], g[4])
	g.AddEdge(2, g[3], g[4])

	/*
		v1 := graph.NewVertex("A")
		v2 := graph.NewVertex("B")
		v3 := graph.NewVertex("C")
		v4 := graph.NewVertex("D")
		v5 := graph.NewVertex("E")
	*/

	fmt.Println("DepthFirstSearch")
	graph.DepthFirstSearch(g[0])
	g.ResetVisited()
	fmt.Println("BreadthFirstSearch")
	graph.BreadthFirstSearch(g[0])

}
