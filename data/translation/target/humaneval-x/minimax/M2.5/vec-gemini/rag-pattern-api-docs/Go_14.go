func AllPrefixes(str string) []string {
    result := []string{}

    for i := 0; i < len(str); i++ {
        result = append(result, str[:i+1])
    }
    return result
}