package main

// PairsSumToZero takes a slice of integers as an input.
// it returns true if there are two distinct elements in the slice that
// sum to zero, and false otherwise.
func PairsSumToZero(l []int) bool {
	for i, l1 := range l {
		for j := i + 1; j < len(l); j++ {
			if l1+l[j] == 0 {
				return true
			}
		}
	}
	return false
}
