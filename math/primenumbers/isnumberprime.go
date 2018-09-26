package main

// O(sqrt(n))

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if isprime(n) {
		fmt.Println(n, "is prime")
	} else {
		fmt.Println(n, "is not prime")
	}
}

func isprime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
