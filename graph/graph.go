package graph

// a tree is a type of graph but not all graphs are trees
//A tree us a connected graph without cycles
// Graphs can be directed or undrected.

//If there us a path between every pair of vertices, is called connected

type Vertex struct {
	value     string
	visited   bool
	neighbors []*Vertex
}

type Graph []*Vertex

func (g *Graph) AddVertex(values ...string) {
	for _, val := range values {
		vertex := Vertex{value: val, visited: false, neighbors: nil}
		*g = append(*g, &vertex)
	}
}

func (g Graph) AddEdge(index int, vertex ...*Vertex) {
	g[index].neighbors = append(g[index].neighbors, vertex...)
}

func (g Graph) ResetVisited() {
	for i, _ := range g {
		g[i].visited = false
	}
}
