func IsNested(s string) bool {
	var openingIndices []int
	var closingIndices []int

	for i, c := range s {
		if c == '[' {
			openingIndices = append(openingIndices, i)
		} else {
			closingIndices = append(closingIndices, i)
		}
	}

	// Reverse the closing indices
	for i, j := 0, len(closingIndices)-1; i < j; i, j = i+1, j-1 {
		closingIndices[i], closingIndices[j] = closingIndices[j], closingIndices[i]
	}

	count := 0
	i := 0
	l := len(closingIndices)

	for _, idx := range openingIndices {
		if i < l && idx < closingIndices[i] {
			count++
			i++
		}
	}

	return count >= 2
}
