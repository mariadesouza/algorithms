package main

import (
	"fmt"

	"github.com/mariadesouza/algorithms/trees/binarysearchtree"
)

func main() {
	root := binarysearchtree.NewNode(5)
	binarysearchtree.Insert(root, 3)
	binarysearchtree.Insert(root, 8)
	binarysearchtree.Insert(root, 12)
	binarysearchtree.Insert(root, 10)

	fmt.Println("InOrder")
	binarysearchtree.InOrder(root)
	fmt.Println("PreOrder")
	binarysearchtree.PreOrder(root)
	fmt.Println("PostOrder")
	binarysearchtree.PostOrder(root)
	SearchTree(root, 13)
	SearchTree(root, 12)
}

func SearchTree(root *binarysearchtree.Node, d int) {
	n := binarysearchtree.Search(root, d)
	if n == nil {
		fmt.Println("Not Found ", d)
	} else {
		fmt.Println("Found ", n.Data)
	}
}
