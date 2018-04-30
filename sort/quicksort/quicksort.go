package main

import "fmt"

/*

Quicksort
Divide and conquer algorithm

Average complexity:  O(nlogn)
Worst case O(n2) but very reare

1. Pick a pivot element
2. Partition based on pivot
3. recursively apply 1 and 2


*/

func main() {
	s := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}

	fmt.Println(quicksort(s))

}

func quicksort(a []int) []int {
	n := len(a)
	pindex := int(n / 2)
	pivot := a[pindex]

	var rightArr, leftArr, equalArr []int
	for i := 0; i < n; i++ {
		if a[i] < pivot {
			leftArr = append(leftArr, a[i])
		} else if a[i] > pivot {
			rightArr = append(rightArr, a[i])
		} else {
			equalArr = append(equalArr, a[i])
		}
	}

	if len(leftArr) > 1 {
		leftArr = quicksort(leftArr)
	}

	if len(rightArr) > 1 {
		rightArr = quicksort(rightArr)
	}

	resultArr := append(leftArr, equalArr...)
	resultArr = append(resultArr, rightArr...)

	return resultArr
}
