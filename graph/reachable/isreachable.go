package main

type Vertex struct {
	Value     string
	Visited   bool
	Neighbors []*Vertex
}

type Graph []*Vertex

func (g *Graph) isReachable(v1, v2 *Vertex) bool {
	if v1 == v2 {
		return true
	}

	return false
}

func main() {

}
