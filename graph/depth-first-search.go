package graph

import "fmt"

// DepthFirstSearch : Start at root abd explore each branch completely before moving to the next branch
//DFS is recursive
func DepthFirstSearch(v *Vertex) {
	if v.visited { // return if visited
		return
	}
	//print node
	v.visited = true
	fmt.Println(v.value)
	for _, vert := range v.neighbors { //visit all neighbors recursively
		DepthFirstSearch(vert)
	}

}
