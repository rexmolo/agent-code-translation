package main

import "fmt"

func FixSpaces(text string) string {
	newText := ""
	i := 0
	start, end := 0, 0
	for i < len(text) {
		if text[i] == ' ' {
			end++
		} else {
			if end-start > 2 {
				newText += "-" + string(text[i])
			} else if end-start > 0 {
				newText += "_" + string(text[i])
			} else {
				newText += string(text[i])
			}
			start, end = i+1, i+1
		}
		i++
	}
	if end-start > 2 {
		newText += "-"
	} else if end-start > 0 {
		newText += "_"
	}
	return newText
}

func main() {
	fmt.Println(FixSpaces("Example"))      // Expected: Example
	fmt.Println(FixSpaces("Example 1"))    // Expected: Example_1
	fmt.Println(FixSpaces(" Example 2"))  // Expected: _Example_2
	fmt.Println(FixSpaces(" Example   3")) // Expected: _Example-3
}