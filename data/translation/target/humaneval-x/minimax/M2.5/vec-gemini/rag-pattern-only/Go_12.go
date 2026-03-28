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