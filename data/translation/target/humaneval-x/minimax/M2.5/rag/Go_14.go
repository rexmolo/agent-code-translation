func AllPrefixes(str string) []string {
	result := make([]string, 0, len(str))

	for i := range str {
		result = append(result, str[:i+1])
	}
	return result
}