package main

import "slices"

func Encrypt(s string) string {
	d := "abcdefghijklmnopqrstuvwxyz"
	out := make([]byte, 0, len(s))
	for _, c := range s {
		idx := slices.Index([]byte(d), byte(c))
		if idx != -1 {
			newIdx := (idx + 2*2) % 26
			out = append(out, d[newIdx])
		} else {
			out = append(out, byte(c))
		}
	}
	return string(out)
}

func main() {
	// Test cases
	println(Encrypt("hi") + " should be lm")
	println(Encrypt("asdfghjkl") + " should be ewhjklnop")
	println(Encrypt("gf") + " should be kj")
	println(Encrypt("et") + " should be ix")
}
