package main

type Node struct {
	Value       int
	Left, Right *Node
}

func New(value int) *Node {
	return &Node{Value: value, Left: nil, Right: nil}
}

/*

An empty tree is height-balanced. A non-empty binary tree T is balanced if:
1) Left subtree of T is balanced
2) Right subtree of T is balanced
3) The difference between heights of left subtree and right subtree is not more than 1.

*/

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func height(root *Node) int {
	if root == nil {
		return 0
	}

	return 1 + max(height(root.Left), height(root.Right))
}

func isBalanced(root *Node) int {
	if root == nil {
		return 1
	}

	var lh, rh int
	lh = height(root.Left)
	rh = height(root.Right)

	if abs(lh-rh) <= 1 && isBalanced(root.Left) > 0 && isBalanced(root.Right) > 0 {
		return 1
	}
	return 0
}

func main() {

}
