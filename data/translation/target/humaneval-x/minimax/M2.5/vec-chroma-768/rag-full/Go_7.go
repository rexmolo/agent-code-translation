func FilterBySubstring(stringList []string, substring string) []string {
	var result []string
	for _, s := range stringList {
		if strings.Contains(s, substring) {
			result = append(result, s)
		}
	}
	return result
}