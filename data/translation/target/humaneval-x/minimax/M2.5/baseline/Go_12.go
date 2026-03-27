package main

func Longest(strings []string) interface{} {
	if len(strings) == 0 {
		return nil
	}

	maxLen := 0
	for _, s := range strings {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}

	for _, s := range strings {
		if len(s) == maxLen {
			return s
		}
	}

	return nil
}

func main() {
	// Test cases
	var result1 interface{} = Longest([]string{})
	println(result1) // <nil>

	var result2 interface{} = Longest([]string{"a", "b", "c"})
	println(result2) // a

	var result3 interface{} = Longest([]string{"a", "bb", "ccc"})
	println(result3) // ccc
}
