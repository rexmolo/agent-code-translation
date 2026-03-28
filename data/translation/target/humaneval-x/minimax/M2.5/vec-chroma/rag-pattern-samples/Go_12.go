package main

func Longest(strings []string) interface{} {
	if len(strings) == 0 {
		return nil
	}

	maxlen := 0
	for _, s := range strings {
		if len(s) > maxlen {
			maxlen = len(s)
		}
	}

	for _, s := range strings {
		if len(s) == maxlen {
			return s
		}
	}

	return nil
}