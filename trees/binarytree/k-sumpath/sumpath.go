package main

import "fmt"

/* Find all k-sum paths in a binary tree
1
| \
7  3
  / \
  4  5


1 7
1 3 4
3 5
*/

type Node struct {
	Item        int
	Left, Right *Node
}

func NewNode(v int) *Node {
	return &Node{Item: v, Left: nil, Right: nil}
}

func main() {
	root := NewNode(1)
	root.Left = NewNode(7)
	root.Right = NewNode(3)
	rr := root.Right
	rr.Left = NewNode(5)
	rr.Right = NewNode(4)

	var path []int
	printKSumPath(root, path, 8)
}

// Preorder traversal (node->left->right)
// we need a slice (vector) to store the path

func printKSumPath(r *Node, path []int, k int) {
	if r == nil {
		return
	}

	// add node
	path = append(path, r.Item)

	//traverse left
	printKSumPath(r.Left, path, k)

	// traverse right
	printKSumPath(r.Right, path, k)

	// check if a k-sum path stops at this node
	sum := 0
	for i := len(path) - 1; i >= 0; i-- {
		sum += path[i]
		if sum == k {
			printVector(path, i)
		}
	}
	// remove the current element from the path
	path = path[:len(path)-1]
}

func printVector(path []int, i int) {
	for j := i; j < len(path); j++ {
		fmt.Print(path[j], " ")
	}
	fmt.Println()
}
