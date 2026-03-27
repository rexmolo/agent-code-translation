package main

func CanArrange(arr []int) int {
	ind := -1
	i := 1
	for i < len(arr) {
		if arr[i] < arr[i-1] {
			ind = i
		}
		i++
	}
	return ind
}