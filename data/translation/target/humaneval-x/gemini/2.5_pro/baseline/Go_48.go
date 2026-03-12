package main

import "fmt"

// IsPalindrome checks if a given string is a palindrome.
// It correctly handles Unicode characters.
func IsPalindrome(text string) bool {
	// To handle multi-byte characters (like emojis or other symbols) correctly,
	// we convert the string to a slice of runes.
	runes := []rune(text)

	// We use a classic two-pointer approach to check for palindromes.
	// 'i' starts at the beginning, 'j' starts at the end.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// If the runes at the pointers do not match, it's not a palindrome.
		if runes[i] != runes[j] {
			return false
		}
	}

	// If the loop completes without finding any mismatches, the string is a palindrome.
	return true
}

// main function to demonstrate IsPalindrome and match Python's doctests.
func main() {
	fmt.Println(IsPalindrome(""))       // Python's >>> is_palindrome('') -> True
	fmt.Println(IsPalindrome("aba"))     // Python's >>> is_palindrome('aba') -> True
	fmt.Println(IsPalindrome("aaaaa"))   // Python's >>> is_palindrome('aaaaa') -> True
	fmt.Println(IsPalindrome("zbcd"))    // Python's >>> is_palindrome('zbcd') -> False
}
