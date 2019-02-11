package graph

import "fmt"

// BreadthFirstSearch : start at root and explore each neighbor before going ot the children
// USED TO find shortest path between any two nodes
// BFS is NOT recursive
//It uses a queue
//O(k^d) => k is no of nodes d is depth
func BreadthFirstSearch(v *Vertex) {

	var q []*Vertex

	v.visited = true
	q = append(q, v) //enqueue

	for len(q) > 0 {

		vert := q[0] //front
		fmt.Println(vert.value)
		q = q[1:] //dequeue

		for _, i := range vert.neighbors {
			if i.visited {
				continue
			}
			i.visited = true
			q = append(q, i) //enqueue
		}
	}
}
