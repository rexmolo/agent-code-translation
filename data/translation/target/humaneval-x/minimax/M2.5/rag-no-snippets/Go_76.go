package main

func IsSimplePower(x int, n int) bool {
	if n == 1 {
		return x == 1
	}
	power := 1
	for power < x {
		power = power * n
	}
	return power == x
}