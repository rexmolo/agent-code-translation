package main

import "fmt"

func FindMax(words []string) string {
	if len(words) == 0 {
		return ""
	}

	result := words[0]
	maxUnique := countUnique(words[0])

	for i := 1; i < len(words); i++ {
		currentUnique := countUnique(words[i])

		if currentUnique > maxUnique {
			result = words[i]
			maxUnique = currentUnique
		} else if currentUnique == maxUnique && words[i] < result {
			result = words[i]
		}
	}

	return result
}

func countUnique(s string) int {
	unique := make(map[rune]bool)
	for _, r := range s {
		unique[r] = true
	}
	return len(unique)
}

func main() {
	// Test cases
	fmt.Println(FindMax([]string{"name", "of", "string"})) // "string"
	fmt.Println(FindMax([]string{"name", "enam", "game"})) // "enam"
	fmt.Println(FindMax([]string{"aaaaaaa", "bb", "cc"}))   // "aaaaaaa"
}