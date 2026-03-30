package main

func CorrectBracketing(brackets string) bool {
	depth := 0
	for _, b := range brackets {
		if b == '(' {
			depth++
		} else {
			depth--
		}
		if depth < 0 {
			return false
		}
	}
	return depth == 0
}

func main() {
	// Test cases
	println(CorrectBracketing("("))      // false
	println(CorrectBracketing("()"))     // true
	println(CorrectBracketing("(()())")) // true
	println(CorrectBracketing(")(()"))  // false
}
