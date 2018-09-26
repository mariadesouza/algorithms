package main

// O(n) with memoization

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("Fibaonacci till %d\n", n)

	if n < 0 {
		os.Exit(0)
	}
	memo := make([]int, n+1)
	fibonacci(n, memo)

	fmt.Println(memo)
}

func fibonacci(n int, memo []int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		memo[n] = n
	} else if memo[n] == 0 {
		memo[n] = fibonacci(n-1, memo) + fibonacci(n-2, memo)
	}
	return memo[n]

}
