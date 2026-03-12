package main

import "fmt"

// SmallestChange calculates the minimum number of elements that need to be changed
// to make the array palindromic. A palindromic array is an array that
// is read the same backwards and forwards. In one change, you can change one
// element to any other element.
//
// For example:
// SmallestChange([]int{1, 2, 3, 5, 4, 7, 9, 6}) == 4
// SmallestChange([]int{1, 2, 3, 4, 3, 2, 2}) == 1
// SmallestChange([]int{1, 2, 3, 2, 1}) == 0
func SmallestChange(arr []int) int {
	ans := 0
	n := len(arr)
	for i := 0; i < n/2; i++ {
		if arr[i] != arr[n-1-i] {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(SmallestChange([]int{1, 2, 3, 5, 4, 7, 9, 6}))
	fmt.Println(SmallestChange([]int{1, 2, 3, 4, 3, 2, 2}))
	fmt.Println(SmallestChange([]int{1, 2, 3, 2, 1}))
}
