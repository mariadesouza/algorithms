/*
leetcode Unique Path I

A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time.
The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
How many possible unique paths are there?

*/

package main

import (
	"fmt"
)

func main() {
	var row, col int
	fmt.Scan(&row)
	fmt.Scan(&col)

	res := uniquePaths(row, col)
	fmt.Println(res)
}

func uniquePaths(row, col int) int {

	//Init the matrix
	matrix := make([][]int, row)
	for m := range matrix {
		matrix[m] = make([]int, col)
	}

	// Calclulate paths
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 || j == 0 { //edge squares set to 1
				matrix[i][j] = 1
			} else {
				matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
			}
		}
	}

	return matrix[row-1][col-1]

}
