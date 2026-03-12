package main

// ChooseNum takes two positive numbers x and y and returns the
// biggest even integer number that is in the range [x, y] inclusive. If 
// there's no such number, then the function should return -1.
//
// For example:
// ChooseNum(12, 15) = 14
// ChooseNum(13, 12) = -1
func ChooseNum(x, y int) int {
	// First, handle the invalid range where the start is greater than the end.
	if x > y {
		return -1
	}

	// If the upper bound 'y' is even, it must be the largest even number in the range.
	if y%2 == 0 {
		return y
	}

	// If 'y' is odd, the largest possible even number is 'y-1'.
	// We must check if 'y-1' is still within the valid range [x, y].
	// Since we know 'y-1 < y', we only need to check if 'y-1 >= x'.
	if y-1 >= x {
		return y - 1
	}

	// If we reach here, it means 'y' is odd and 'y-1 < x'.
	// This occurs when x and y are the same odd number (e.g., x=13, y=13).
	// In this case, no even number exists in the range.
	return -1
}
