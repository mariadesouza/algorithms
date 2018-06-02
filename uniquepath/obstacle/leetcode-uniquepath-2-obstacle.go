/*
leetcode Unique Path II

Follow up for "Unique Paths":
Now consider if some obstacles are added to the grids. How many unique paths would there be?
An obstacle and empty space is marked as 1 and 0 respectively in the grid.
For example,
There is one obstacle in the middle of a 3x3 grid as illustrated below.
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
The total number of unique paths is 2.


Analysis:
Just a little bit changes from the previous problem.
Note that if the start position is obstacle, then return 0.
Pretty easy if you solve the previous one.

*/

package main

import (
	"fmt"
)

func main() {
	obstacleMatrix := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	res := uniquePathWithObstacle(a)
	fmt.Println(res)
}

func uniquePathWithObstacle(obstacleMatrix [][]int) int {

	if len(obstacleMatrix) == 0 {
		return 0
	}

	//if the start position is obstacle, then return 0.
	if obstacleMatrix[0][0] == 1 {
		return 0
	}

	row, col := len(obstacleMatrix), len(obstacleMatrix[0])

	//Init Path matrix
	matrix := make([][]int, row)
	for i := range matrix {
		matrix[i] = make([]int, col)
	}

	// Calclulate paths
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {

			if obstacleMatrix[i][j] == 1 {
				continue
			}

			if i == 0 && j == 0 {
				matrix[0][0] = 1
			} else {
				if i == 0 {
					matrix[0][j] = matrix[0][j-1]
				} else if j == 0 { //edge squares
					matrix[i][0] = matrix[i-1][0]
				} else {
					matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
				}
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	return matrix[row-1][col-1]

}
