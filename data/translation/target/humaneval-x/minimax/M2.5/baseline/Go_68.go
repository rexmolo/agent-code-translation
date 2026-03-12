func Pluck(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	// Find smallest even value and its first occurrence index
	smallestEven := -1
	index := -1

	for i, v := range arr {
		if v%2 == 0 {
			if smallestEven == -1 || v < smallestEven {
				smallestEven = v
				index = i
			}
		}
	}

	if index == -1 {
		return []int{}
	}

	return []int{smallestEven, index}
}