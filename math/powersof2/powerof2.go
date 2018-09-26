package main

// O(log n)
// n = 50
// power(25)
//	power(12)
//		power(6)
//			power(3)
//			 power (1)
//			 print & return 1
//		print & return 2
//     print & return 4
//  print & return 8
// print & return 16
// print & return 32

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Powers of 2 till n")
	var n int
	fmt.Scan(&n)
	if n < 0 {
		os.Exit(0)
	}
	powersOf2(n)
}

func powersOf2(n int) int {
	if n <= 1 {
		fmt.Println(n)
		return n
	}
	prev := powersOf2(n / 2)
	curr := prev * 2
	fmt.Println(curr)
	return curr

}
