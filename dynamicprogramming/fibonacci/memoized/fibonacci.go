package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)
	memo := make([]int64, n)
	fmt.Println(fibonacci(n, memo))
}

func fibonacci(n int64, memo []int64) int64 {

	// already computed  - return
	if memo[n-1] != 0 {
		return memo[n-1]
	}

	var result int64
	if n == 1 || n == 2 {
		result = 1
	} else {
		result = fibonacci(n-1, memo) + fibonacci(n-2, memo)
	}
	memo[n-1] = result
	return result
}
