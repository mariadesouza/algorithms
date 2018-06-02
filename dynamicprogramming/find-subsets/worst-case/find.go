/*
  find subsets that add to a certain number
*/
package main

import "fmt"

func main() {
	var findSum int
	fmt.Scan(&findSum)
	a := []int{2, 4, 6, 10}
	fmt.Println(countSubsets(a, findSum, len(a)-1))
}

func countSubsets(a []int, findSum int, i int) int {
	// if sum is 0 return 0
	if findSum == 0 {
		return 1
	}
	if findSum < 0 || i < 0 {
		return 0
	}

	if findSum < a[i] {
		return countSubsets(a, findSum, i-1)
	} else {
		return countSubsets(a, findSum-a[i], i-1) + countSubsets(a, findSum, i-1)
	}

}
