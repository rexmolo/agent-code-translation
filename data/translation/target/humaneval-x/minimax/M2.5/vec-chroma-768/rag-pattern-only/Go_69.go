func Search(lst []int) int {
	// Find the maximum value in the list
	maxVal := lst[0]
	for _, v := range lst {
		if v > maxVal {
			maxVal = v
		}
	}

	// Create frequency array where index represents the integer value
	frq := make([]int, maxVal+1)
	for _, i := range lst {
		frq[i]++
	}

	// Find the greatest integer i such that frequency >= i
	ans := -1
	for i := 1; i < len(frq); i++ {
		if frq[i] >= i {
			ans = i
		}
	}

	return ans
}