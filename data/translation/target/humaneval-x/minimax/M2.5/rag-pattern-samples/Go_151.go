package main

func DoubleTheDifference(lst []float64) int {
	sum := 0
	for _, i := range lst {
		// Skip non-positive numbers
		if i <= 0 {
			continue
		}
		// Check if it's an integer (no decimal part)
		// In Python: "." not in str(i)
		if i != float64(int64(i)) {
			continue
		}
		// Check if odd: in Python i%2 != 0
		if int64(i)%2 == 1 {
			square := int(i * i)
			sum += square
		}
	}
	return sum
}