package main

import "strings"

func Encode(message string) string {
	vowels := "aeiouAEIOU"

	// Create replacement map for vowels (a->c, e->g, i->k, o->q, u->w)
	vowelsReplace := make(map[rune]rune, len(vowels))
	for _, v := range vowels {
		vowelsReplace[v] = rune(v) + 2
	}

	// Convert to rune slice for processing
	var result []rune
	for _, ch := range message {
		// Swap case: subtract 32 from lowercase to get uppercase,
		// add 32 to uppercase to get lowercase (ASCII trick)
		if ch >= 'a' && ch <= 'z' {
			ch -= 32
		} else if ch >= 'A' && ch <= 'Z' {
			ch += 32
		}

		// Replace vowel if in map
		if repl, ok := vowelsReplace[ch]; ok {
			ch = repl
		}
		result = append(result, ch)
	}

	return string(result)
}
