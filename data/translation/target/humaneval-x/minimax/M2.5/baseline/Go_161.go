package main

import "fmt"

func Solve(s string) string {
	flg := 0
	idx := 0
	newStr := []rune(s)

	for _, r := range s {
		if isAlpha(r) {
			newStr[idx] = swapCase(r)
			flg = 1
		}
		idx++
	}

	result := string(newStr)

	if flg == 0 {
		// Reverse the string
		runes := []rune(result)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	return result
}

// isAlpha checks if a rune is an alphabetic character
func isAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

// swapCase converts a letter from lower to upper or vice versa
func swapCase(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 'a' + 'A'
	}
	if r >= 'A' && r <= 'Z' {
		return r - 'A' + 'a'
	}
	return r
}

func main() {
	fmt.Println(Solve("1234")) // 4321
	fmt.Println(Solve("ab"))    // AB
	fmt.Println(Solve("#a@C")) // #A@c
}
