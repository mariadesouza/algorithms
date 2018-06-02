/*
 binary search, also known as half-interval search, logarithmic search, or binary chop, is a search algorithm that
 finds the position of a target value within a sorted array.

 Binary search runs in at worst logarithmic time, making O(log n) comparisons,

 Binary search works on sorted arrays

 compare the middle element of the array with the target value

 divide and search based on value

 */

package main

import (
  "fmt"
  "os"
)

func main(){
  sortedArr := []int{2, 5, 7, 11, 23, 32, 36}

  fmt.Println(sortedArr)
  fmt.Println("enter number to search for:")
	var target int
  fmt.Scanf("%d",&target)

  low := 0
  high := len(sortedArr) -1

  //TODO: return gracefully
  if (high < low){
    fmt.Println("Empty array")
    os.Exit(0)
  }

  result := binary_search(sortedArr, low, high, target)

  if (result == -1){
    fmt.Println("not found")
  } else {
    fmt.Println("Found at index: ", result)
  }
  os.Exit(1)
}


func binary_search(a []int, low, high, target int) int{

  if (low == high){
    if a[low] == target{
      fmt.Println("target found")
      return low
      } else {
          fmt.Println("no target found")
        return -1
      }
  }
  middle := int((low + high)/2)
  if a[middle] == target{
    return middle
  }  else if target < a[middle]{
    return binary_search(a, low, middle-1, target)
  } else {
    return binary_search(a, middle+1, high, target)
  }
}
