package main

func CountUpTo(n int) []int {
	primes := []int{}

	for i := 2; i < n; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}

func main() {
	// Test cases
	println(CountUpTo(5))
	println(CountUpTo(11))
	println(CountUpTo(0))
	println(CountUpTo(20))
	println(CountUpTo(1))
	println(CountUpTo(18))
}
