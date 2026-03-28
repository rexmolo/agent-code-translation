func Encrypt(s string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	result := ""
	for _, c := range s {
		idx := strings.Index(alphabet, string(c))
		if idx != -1 {
			result += string(alphabet[(idx+2*2)%26])
		} else {
			result += string(c)
		}
	}
	return result
}