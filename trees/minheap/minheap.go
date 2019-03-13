package main

//complete binary tree (totally filled other than the rightmost element on the last level)\
//Each node in a min-heap is smaller than its children
// The root is the minimum elemnent in the tree

// KEY OPERATIONS
//- insert
// - extract_min

//INSERT O(logn)
//1.
// alwyas start inserting at the bottom
// insert at the rightmost spot to mantain ccomplete tree property
//2.
//Fix the tree by swapping the new element with its parent until we find an appropriate spot
// Bubble up the minimum element

/*
	 2
  /    \
 50     4
 /\    /\
55 90 87 7
*/

//EXtract Minimum Element
// Minimum element always at the top
// swap witht he bottom-most rightmost element
// and then bubble it down, swapping it with one of its childrem

func main() {

}
