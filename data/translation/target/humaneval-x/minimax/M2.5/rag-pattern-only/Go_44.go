func ChangeBase(x int, base int) string {
	if x == 0 {
		return ""
	}

	var digits []rune
	for x > 0 {
		digit := x % base
		digits = append(digits, rune('0'+digit))
		x /= base
	}

	// Reverse the digits to get correct order
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return string(digits)
}