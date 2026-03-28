package main

import "fmt"

func encodeCyclic(s string) string {
	// split string to groups. Each of length 3.
	numGroups := (len(s) + 2) / 3
	var groups []string
	for i := 0; i < numGroups; i++ {
		start := 3 * i
		end := start + 3
		if end > len(s) {
			end = len(s)
		}
		groups = append(groups, s[start:end])
	}
	// cycle elements in each group. Unless group has fewer elements than 3.
	for i, group := range groups {
		if len(group) == 3 {
			groups[i] = group[1:] + group[0:1]
		}
	}
	result := ""
	for _, group := range groups {
		result += group
	}
	return result
}

func DecodeCyclic(s string) string {
	return encodeCyclic(encodeCyclic(s))
}

func main() {
	fmt.Println(DecodeCyclic("abcdef"))   // Output: abcdef
	fmt.Println(DecodeCyclic("bcaefd"))   // Output: abcdef
	fmt.Println(DecodeCyclic("abc"))      // Output: abc
	fmt.Println(DecodeCyclic("ab"))       // Output: ab
	fmt.Println(DecodeCyclic("a"))        // Output: a
}
