package main

import (
	"fmt"
	"strconv"
)

func EvenOddCount(num int) [2]int {
 evenCount := 0
 oddCount := 0
 
 // Handle negative numbers by working with absolute value
 absNum := num
 if absNum < 0 {
  absNum = -absNum
 }
 
 // Convert to string to iterate over digits
 strNum := strconv.Itoa(absNum)
 
 // Iterate over each character (digit)
 for _, c := range strNum {
  digit := int(c - '0') // Convert rune to integer digit
  if digit%2 == 0 {
   evenCount++
  } else {
   oddCount++
  }
 }
 
 return [2]int{evenCount, oddCount}
}

func main() {
	// Test the function
	fmt.Println(EvenOddCount(-12)) // Output: [1 1]
	fmt.Println(EvenOddCount(123)) // Output: [1 2]
}