package main

func HowManyTimes(str string, substring string) int {
	times := 0
	subLen := len(substring)

	if subLen == 0 || len(str) == 0 {
		return 0
	}

	for i := 0; i <= len(str)-subLen; i++ {
		if str[i:i+subLen] == substring {
			times++
		}
	}

	return times
}
