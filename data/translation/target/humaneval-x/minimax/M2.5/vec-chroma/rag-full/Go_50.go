package main

func DecodeShift(s string) string {
	result := make([]rune, 0, len(s))
	for _, ch := range s {
		decoded := (int(ch) - int('a') - 5 + 26) % 26 + int('a')
		result = append(result, rune(decoded))
	}
	return string(result)
}
