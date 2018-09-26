package main

import (
	"fmt"
)

// O(n)

func main() {
	fmt.Println("Factorial of?")
	var n int
	fmt.Scan(&n)
	res := factorial(n)
	fmt.Println(res)
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	result := n * factorial(n-1)
	return result
}
