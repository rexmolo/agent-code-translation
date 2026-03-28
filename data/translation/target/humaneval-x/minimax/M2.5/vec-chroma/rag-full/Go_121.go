package main

func Solution(lst []int) int {
    sum := 0
    for idx, x := range lst {
        if idx%2 == 0 && x%2 == 1 {
            sum += x
        }
    }
    return sum
}

// Example usage:
func main() {
    // Test examples
    result1 := Solution([]int{5, 8, 7, 1})
    result2 := Solution([]int{3, 3, 3, 3, 3})
    result3 := Solution([]int{30, 13, 24, 321})
    
    // Expected: 12, 9, 0
    println(result1, result2, result3)
}
