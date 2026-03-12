package main

func SeparateParenGroups(parenString string) []string {
	result := []string{}
	currentString := []byte{}
	currentDepth := 0

	for i := 0; i < len(parenString); i++ {
		c := parenString[i]
		if c == '(' {
			currentDepth++
			currentString = append(currentString, c)
		} else if c == ')' {
			currentDepth--
			currentString = append(currentString, c)

			if currentDepth == 0 {
				result = append(result, string(currentString))
				currentString = currentString[:0] // Clear the slice
			}
		}
		// Ignore spaces and any other characters
	}

	return result
}
