package main

func StringXor(a string, b string) string {
	// Determine the shorter length (zip behavior in Python stops at shorter string)
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}

	// Create result slice
	result := make([]byte, minLen)

	for i := 0; i < minLen; i++ {
		if a[i] == b[i] {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}

	return string(result)
}
