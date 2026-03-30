func CountUpper(s string) int {
	count := 0
	for i := 0; i < len(s); i += 2 {
		c := s[i]
		if c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
			count++
		}
	}
	return count
}