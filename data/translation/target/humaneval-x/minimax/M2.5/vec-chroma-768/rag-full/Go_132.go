package main

func isNested(s string) bool {
	var openingIndices []int
	var closingIndices []int

	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			openingIndices = append(openingIndices, i)
		} else {
			closingIndices = append(closingIndices, i)
		}
	}

	// Reverse closing indices
	for i, j := 0, len(closingIndices)-1; i < j; i, j = i+1, j-1 {
		closingIndices[i], closingIndices[j] = closingIndices[j], closingIndices[i]
	}

	cnt := 0
	i := 0
	l := len(closingIndices)

	for _, idx := range openingIndices {
		if i < l && idx < closingIndices[i] {
			cnt++
			i++
		}
	}

	return cnt >= 2
}

func main() {
	// Test cases
	println(isNested("[[]]"))        // true
	println(isNested("[]]]]]]][[[[[]")) // false
	println(isNested("[][]"))        // false
	println(isNested("[]"))           // false
	println(isNested("[[][]]"))      // true
	println(isNested("[[]][["))       // true
}
