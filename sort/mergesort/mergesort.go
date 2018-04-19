/*
Merge Sort
It divides iput to two halves and then merges the sorted halves

1. Find middle point to divide the array
    middle = (l+ r) / 2
2. mergesort(:middle) // first half
3. mergesort(middle:) //second half
4. merge arrays from 2 and 3

Pro : Complexity Always O(nlogn)
Needs extra space for merging
*/

package main

import (
	"fmt"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	s := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	fmt.Println(s)
	a := mergeSort(s)
	fmt.Println(a)

}

///recursive
func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	middle := len(a) / 2
	l := mergeSort(a[:middle])
	r := mergeSort(a[middle:])
	return merge(l, r)
}

func merge(l, r []int) []int {
	a := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(a, r...)
		}
		if len(r) == 0 {
			return append(a, l...)
		}
		if l[0] <= r[0] {
			a = append(a, l[0])
			l = l[1:]

		} else {
			a = append(a, r[0])
			r = r[1:]
		}
	}

	return a
}
