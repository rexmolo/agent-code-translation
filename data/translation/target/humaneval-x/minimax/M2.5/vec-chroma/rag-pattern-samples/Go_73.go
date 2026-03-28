package main

func SmallestChange(arr []int) int {
	ans := 0
	n := len(arr)
	for i := 0; i < n/2; i++ {
		if arr[i] != arr[n-i-1] {
			ans++
		}
	}
	return ans
}
