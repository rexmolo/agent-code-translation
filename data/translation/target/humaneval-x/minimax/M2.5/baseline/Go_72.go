package main

func WillItFly(q []int, w int) bool {
	// Check if the sum of elements is greater than the maximum weight
	sum := 0
	for _, v := range q {
		sum += v
	}
	if sum > w {
		return false
	}

	// Check if the list is a palindrome (balanced)
	i, j := 0, len(q)-1
	for i < j {
		if q[i] != q[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// For testing purposes
func main() {
	// Test cases
	println(WillItFly([]int{1, 2}, 5))   // false
	println(WillItFly([]int{3, 2, 3}, 1)) // false
	println(WillItFly([]int{3, 2, 3}, 9)) // true
	println(WillItFly([]int{3}, 5))       // true
}
