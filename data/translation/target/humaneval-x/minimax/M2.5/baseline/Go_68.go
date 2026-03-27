package main

func Pluck(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	minVal := -1
	minIdx := -1

	for i, v := range arr {
		if v%2 == 0 {
			if minVal == -1 || v < minVal {
				minVal = v
				minIdx = i
			}
		}
	}

	if minIdx == -1 {
		return []int{}
	}

	return []int{minVal, minIdx}
}
