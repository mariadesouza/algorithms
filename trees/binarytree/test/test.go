package main

import (
	"fmt"

	"github.com/mariadesouza/algorithms/trees/binarytree"
)

func main() {

	// If it is just a Binary tree we dont need to follow the rule left < current < right
	b := binarytree.NewNode(5)
	b.Left = binarytree.NewNode(8)
	b.Left.Left = binarytree.NewNode(10)
	b.Right = binarytree.NewNode(3)
	print(b)
}

func print(b *binarytree.BinaryTreeNode) {
	if b == nil {
		return
	}
	if b.Left != nil {
		print(b.Left)
	}
	if b.Right != nil {
		print(b.Right)
	}
	fmt.Println(b.Data)

}
