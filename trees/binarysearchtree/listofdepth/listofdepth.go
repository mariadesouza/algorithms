package main

import (
	"container/list"
	"fmt"
)

/*
Level Order Traversal

given a binary tree,
design an algorithm to create a linked list of all nodes at each depth
E.g. If you have a tree with depth D, you'll have D linked lists
Input Tree
       A
      / \
     B   C
    / \   \
   D   E   F

Output Tree
Input Tree
       A--> null
      / \
     B --> C--> null
    / \   \
   D --> E-->F--> null

*/

type Node struct {
	Item        int
	Left, Right *Node
}

func NewNode(i int) *Node {
	return &Node{Item: i, Left: nil, Right: nil}
}

func main() {

	sortedItems := []int{2, 5, 8, 14, 35, 42, 64, 73, 85, 92}
	root := makeBinarySearchTree(sortedItems)

	var l []*list.List
	l = createLevelLinkedLists(root, l, 0)

	for _, listLevel := range l {
		//	fmt.Println(i, listLevel)
		for e := listLevel.Front(); e != nil; e = e.Next() {
			fmt.Printf("%d -> ", e.Value.(*Node).Item)
		}
		fmt.Println()
	}
	//preorder(root)
}

func createLevelLinkedLists(n *Node, l []*list.List, level int) []*list.List {
	if n == nil {
		return l
	}
	if len(l) == level { //first time at level - create the list and add to slice
		listLevel := list.New()
		l = append(l, listLevel)
	}
	l[level].PushBack(n)
	fmt.Println(level, len(l))
	l = createLevelLinkedLists(n.Left, l, level+1)
	l = createLevelLinkedLists(n.Right, l, level+1)

	return l
}

func makeBinarySearchTree(a []int) *Node {

	if len(a) == 0 {
		return nil
	}

	mid := len(a) / 2
	root := NewNode(a[mid])
	root.Left = makeBinarySearchTree(a[0:mid])
	root.Right = makeBinarySearchTree(a[mid+1 : len(a)])

	return root

}
