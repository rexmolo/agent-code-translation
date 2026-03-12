import (
	"sort"
)

func Median(l []int) float64 {
	// Create a copy of the slice to avoid modifying the original
	sorted := make([]int, len(l))
	copy(sorted, l)

	sort.Ints(sorted)

	n := len(sorted)
	if n%2 == 1 {
		return float64(sorted[n/2])
	}
	return (float64(sorted[n/2-1]) + float64(sorted[n/2])) / 2.0
}