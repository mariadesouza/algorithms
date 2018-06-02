package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n <= 0 {
		os.Exit(0)
	}
	fmt.Println(fibonacci(n))
}

func fibonacci(n int) int64 {
	if n == 1 || n == 2 {
		return 1
	}
	memo := make([]int64, n)
	memo[0] = 1
	memo[1] = 1

	for i := 2; i < n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n-1]
}
