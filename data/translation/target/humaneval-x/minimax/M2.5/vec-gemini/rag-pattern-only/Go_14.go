func AllPrefixes(str string) []string {
	result := make([]string, 0, len(str))

	for i := 0; i < len(str); i++ {
		result = append(result, str[:i+1])
	}
	return result
}