package main

func LargestSmallestIntegers(lst []int) [2]interface{} {
	var largestNegative, smallestPositive int
	hasNegative, hasPositive := false, false

	for _, v := range lst {
		if v < 0 {
			if !hasNegative || v > largestNegative {
				largestNegative = v
				hasNegative = true
			}
		} else if v > 0 {
			if !hasPositive || v < smallestPositive {
				smallestPositive = v
				hasPositive = true
			}
		}
	}

	result := [2]interface{}{}
	if hasNegative {
		result[0] = largestNegative
	}
	if hasPositive {
		result[1] = smallestPositive
	}

	return result
}
