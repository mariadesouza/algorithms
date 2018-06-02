/*

Breaking up a problem into simpler subproblems, solving each of those subproblems just once, and storing their solutions
The important part of dynamic programming is storing the intermediate result set
*/
package main

import "fmt"

func main() {
	var findSum int
	fmt.Scan(&findSum)
	a := []int{2, 4, 6, 10}
	fmt.Println(countSubsets(a, findSum, len(a)-1))
}

func countSubsets(a []int, findSum int, len a) int {

}
