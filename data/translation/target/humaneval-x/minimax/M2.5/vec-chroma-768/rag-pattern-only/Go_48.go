package main

func IsPalindrome(text string) bool {
	runes := []rune(text)
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}
