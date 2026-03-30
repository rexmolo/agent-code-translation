package main

func BelowThreshold(l []int, t int) bool {
	for _, e := range l {
		if e >= t {
			return false
		}
	}
	return true
}