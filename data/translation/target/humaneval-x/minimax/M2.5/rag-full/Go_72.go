package main

func WillItFly(q []int, w int) bool {
	// Check if sum of elements exceeds maximum weight
	sum := 0
	for _, v := range q {
		sum += v
	}
	if sum > w {
		return false
	}

	// Check if the slice is a palindrome (balanced)
	i, j := 0, len(q)-1
	for i < j {
		if q[i] != q[j] {
			return false
		}
		i++
		j--
	}

	return true
}
