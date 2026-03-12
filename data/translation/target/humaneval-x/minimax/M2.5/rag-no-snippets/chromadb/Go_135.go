func CanArrange(arr []int) int {
	ind := -1
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			ind = i
		}
	}
	return ind
}